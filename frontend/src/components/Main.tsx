
import { useState } from "react"
import TaskHolder from "../components/TaskHolder"
import TaskInput from "../components/TaskInput"
import { TaskInputProps } from "../types/type"
import { TaskProps } from "./TaskCard"
import Tabs from "./Tabs"
import { useTasks, useTasksWithUserID } from "../hooks/useTasks"

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
  const [tabs, setTabs] = useState('create')

  const { data, loading, error } = useTasks({ id: 35 })



  const addTask = (task: TaskInputProps) => {
    setTasks([...tasks, task])
  }

  return (
    <div className="flex flex-col justify-center items-center min-h-screen p-6">
      <Tabs activeTab={tabs} setActiveTab={setTabs} />
      {
        tabs === 'create' && <TaskInput addTask={addTask} />
      }
      {
        tabs === 'view' && <TaskHolder tasks={data} />
      }
    </div>
  )
}

export default Main
