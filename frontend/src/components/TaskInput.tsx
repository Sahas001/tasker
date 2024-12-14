import { useState } from "react"
import { TaskInputProps } from "../types/type"

function TaskInput({ addTask }: { addTask: (task: TaskInputProps) => void }) {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [category, setCategory] = useState('')
  const [status, setStatus] = useState('Pending')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (title && description && category) {
      addTask({ title, description, category, status: status == 'Pending' ? false : true, createdAt: new Date().toLocaleDateString() })
    }
    console.log({ title, description, category, status })
  }

  return (
    <div className="flex flex-col items-center space-y-6 p-6 bg-gray-100 min-h-screen">
      <form
        onSubmit={handleSubmit}
        className="w-3/4 bg-white shadow-md rounded-lg p-6 mt-5"
      >
        <h2 className="text-lg font-bold text-gray-800 mb-4">Create a Task</h2>
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="w-full mb-4 p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <textarea
          placeholder="Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="w-full mb-4 p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <input
          type="text"
          placeholder="Category"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
          className="w-full mb-4 p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <select
          value={status}
          onChange={(e) => setStatus(e.target.value)}
          className="w-full mb-4 p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="Pending">Pending</option>
          <option value="Completed">Completed</option>
        </select>
        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600"
        >
          Add Task
        </button>
      </form>
    </div>
  )
}

export default TaskInput
