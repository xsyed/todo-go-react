import React, { useState, useEffect } from 'react';
import TodoForm from './TodoForm';
import Todo from './Todo';

function TodoList() {
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    getTodos()
  },[])

  function getTodos(){
    fetch("http://localhost:9000/todo")
    .then(res => res.json())
    .then((result)=>{
      setTodos(result.data)
    })
  }

  const addTodo = todo => {
    if (!todo.title || /^\s*$/.test(todo.title)) {
      return;
    }

    fetch('http://localhost:9000/todo', {
        method: 'POST',
        body: JSON.stringify({
          title: todo.title 
        }),
        headers: {
          'Content-type': 'application/json; charset=UTF-8',
        },
      })
        .then((response) => response.json())
        .then((json) =>{
          getTodos()
        });
  };

  const updateTodo = (todoId, newValue) => {

    if (!newValue.title || /^\s*$/.test(newValue.title)) {
      return;
    }

    fetch('http://localhost:9000/todo/'+todoId, {
      method: 'PUT',
      body: JSON.stringify({
        id: todoId,
        title: newValue.title,
      }),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
      .then((response) => response.json())
      .then((json) =>{
        getTodos()
      });

    //setTodos(prev => prev.map(item => (item.id === todoId ? newValue : item)));
  };

  const removeTodo = id => {
    const removedArr = [...todos].filter(todo => todo.id !== id);

    fetch('http://localhost:9000/todo/'+id, {
        method: 'DELETE',
        headers: {
          'Content-type': 'application/json; charset=UTF-8',
        },
      })
        .then((response) => response.json())
        .then((json) =>{
          getTodos()
        });

    //setTodos(removedArr);
  };

  const completeTodo = todo => {
   
    fetch('http://localhost:9000/todo/'+todo.id, {
      method: 'PUT',
      body: JSON.stringify({
        id: todo.id,
        title: todo.title,
        completed: !todo.completed
      }),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
      .then((response) => response.json())
      .then((json) =>{
        getTodos()
      });

    
  };

  return (
    <>
      <h1>things to be done</h1>
      <TodoForm onSubmit={addTodo} />
      <Todo
        todos={todos}
        completeTodo={completeTodo}
        removeTodo={removeTodo}
        updateTodo={updateTodo}
      />
    </>
  );
}

export default TodoList;
