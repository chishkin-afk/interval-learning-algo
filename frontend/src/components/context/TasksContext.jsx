import { createContext, useState } from "react";

export const TasksContext = createContext({})

export const TasksProvider = ({children}) => {
    const [tasks, setTasks] = useState([])
    fetch('http://localhost:3001/tasks').then(resp => resp.json()).then(data => { 
        for (const task of data) {
            task.created_at = new Date(task.created_at)
            task.next_notification = new Date(task.next_notification)
        }

        setTasks(data)
    })

    const [newTaskTitle, setNewTaskTitle] = useState('')
    const [newLeetcodeURL, setNewLeetcodeURL] = useState('')

    const addTask = () => {
        console.log('todo add task...')
    }

    return <TasksContext.Provider
        value={{
            tasks,
            newTaskTitle,
            setNewTaskTitle,
            newLeetcodeURL,
            setNewLeetcodeURL,
        }}
    >
        {children}
    </TasksContext.Provider>
}