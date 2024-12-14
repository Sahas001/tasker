type TabsProps = {
  activeTab: string;
  setActiveTab: (tab: string) => void;
}

function Tabs({ activeTab, setActiveTab }: TabsProps) {
  return (
    <div className="flex space-x-4 py-4 px-6 sticky top-0 z-10">
      <button
        onClick={() => setActiveTab('create')}
        className={`px-4 py-2 rounded ${activeTab === 'create' ? 'bg-blue-500 text-white' : 'bg-gray-200'
          } hover:bg-blue-600 hover:text-white`}
      >
        Create Task
      </button>
      <button
        onClick={() => setActiveTab('view')}
        className={`px-4 py-2 rounded ${activeTab === 'view' ? 'bg-blue-500 text-white' : 'bg-gray-200'
          } hover:bg-blue-600 hover:text-white`}
      >
        View Tasks
      </button>
    </div>
  )
}

export default Tabs
