import { createContext } from "react";

export const TasksContext = createContext({})

export const TasksProvider = ({children}) => {
    return <TasksContext.Provider
        value={{}}
    >
        {children}
    </TasksContext.Provider>
}