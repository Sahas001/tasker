
import { useState } from "react"
import TaskHolder from "../components/TaskHolder"
import TaskInput from "../components/TaskInput"
import { TaskInputProps } from "../types/type"
import { TaskProps } from "./TaskCard"

function Main() {
  const [tasks, setTasks] = useState<TaskProps[]>([
    {
      title: 'Task 1',
      description: 'What a long description for Task 1',
      category: 'Work',
      status: true,
      createdAt: '2021-09-01'
    },
    {
      title: 'Task 2',
      description: 'That can not be a long description',
      category: 'Personal',
      status: false,
      createdAt: '2021-09-02'
    },
    {
      title: 'Task 3',
      description: 'This is kind of a long description for Task 3',
      category: 'Work',
      status: true,
      createdAt: '2021-09-03'
    }

  ])

  const addTask = (task: TaskInputProps) => {
    setTasks([...tasks, task])
  }

  return (
    <div className="flex flex-col justify-center items-center p-5 ">
      <TaskInput addTask={addTask} />
      <TaskHolder tasks={tasks} />
    </div>
  )
}

export default Main
