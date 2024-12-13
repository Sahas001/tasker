import TaskCard, { TaskProps } from "./TaskCard"

function TaskHolder({ tasks }: { tasks: TaskProps[] }) {
  return (
    <div className="flex flex-col items-center space-y-6 p-6 bg-gray-100 min-h-screen">
      {tasks.map((task: TaskProps, index) => (
        <TaskCard
          key={index}
          title={task.title}
          description={task.description}
          category={task.category}
          status={task.status}
          createdAt={task.createdAt}
        />
      ))}
    </div>)
}

export default TaskHolder
