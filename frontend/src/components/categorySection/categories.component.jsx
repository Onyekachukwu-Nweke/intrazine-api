const categories = [
  { id: 1, name: 'Business', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', icon: '🏢' },
  { id: 2, name: 'Startup', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', icon: '🚀' },
  { id: 3, name: 'Economy', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', icon: '📈' },
  { id: 4, name: 'Technology', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', icon: '💡' },
];

const CategorySection = () => {
  return (
    <div className="container mx-auto my-16 px-4 text-center">
      <h2 className="text-3xl font-bold mb-8">Choose A Category</h2>
      <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
        {categories.map((category) => (
          <div
            key={category.id}
            className="p-6 border border-gray-300 hover:bg-yellow hover:border-0 transition duration-300 ease-in-out"
          >
            <div className="text-4xl mb-4">{category.icon}</div>
            <h3 className="text-xl font-bold mb-2">{category.name}</h3>
            <p className="text-gray-700">{category.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default CategorySection;
