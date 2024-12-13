
export type TaskProps = {
  title: string;
  description: string;
  category: string;
  status: boolean;
  createdAt: string;
}

function TaskCard({ title, description, category, status, createdAt }: TaskProps) {

  return (
    <>
      <div className="flex items-center justify-center min-h-5 bg-gray-100">
        <div className="max-w-sm w-80 bg-white shadow-md rounded-lg p-6 relative group hover:shadow-lg transition-shadow">
          <h2 className="text-xl font-bold text-gray-800 mb-2">{title}</h2>
          <p className="text-gray-600 mb-4">{description}</p>
          <p className="text-gray-500 mb-4 italic">Category: {category}</p>
          <div className="flex justify-between items-center text-sm text-gray-500">
            <span
              className={`px-3 py-1 rounded-full font-semibold text-sm ${status === true ? 'bg-green-200 text-green-900' : 'bg-yellow-100 text-yellow-800'
                }`}
            >
              {status}
            </span>
            <span>{createdAt}</span>
          </div>
          <div className="absolute inset-0 flex justify-center items-center bg-black bg-opacity-40 opacity-0 group-hover:opacity-100 transition-opacity rounded-lg">
            <button className="text-white bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded mr-4">Update</button>
            <button className="text-white bg-red-500 hover:bg-red-600 px-4 py-2 rounded">Delete</button>
          </div>
        </div>
      </div>
    </>
  )
}

export default TaskCard
