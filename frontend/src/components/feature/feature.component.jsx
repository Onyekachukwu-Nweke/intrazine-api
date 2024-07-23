import featureImage from '../../assets/feature-image.png'; // Ensure this path is correct

const posts = [
  { id: 1, title: '8 Figma design systems that you can download for free today.', author: 'Onyeka', date: 'Aug 2024' },
  { id: 2, title: 'SRE, A Practical Guide', author: 'Onyeka', date: 'Aug 2024' },
  { id: 3, title: 'Daily DevOps Nuggets', author: 'Onyeka', date: 'Aug 2024' },
  { id: 4, title: 'Fullstack Software Engineering, A Lesson Learnt the Hard Way', author: 'Onyeka', date: 'Aug 2024' },
  // Add more posts as needed
];

const FeatureSection = () => {
  return (
    <div className="container mx-auto my-16 px-4">
      <div className="flex flex-col md:flex-row justify-between">
        {/* Feature Post */}
        <div className="md:w-2/3 p-4">
          <h2 className="text-3xl font-bold mb-4">Feature Post</h2>
          <div className="bg-gray-100 p-6 rounded-lg">
            <img src={featureImage} alt="Feature" className="w-full h-64 object-cover rounded-lg mb-4" />
            <p className="text-sm text-gray-600 mb-2">By Onyeka | Aug 2024</p>
            <h3 className="text-2xl font-semibold mb-2">Altschool Africa Experience</h3>
            <p className="mb-4">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus imperdiet, nulla et dictum interdum,
              nisi lorem egestas odio, vitae scelerisque enim ligula venenatis dolor.
            </p>
            <button className="bg-yellow text-white px-6 py-3 rounded hover:bg-gray-800">Read More</button>
          </div>
        </div>

        {/* All Posts */}
        <div className="md:w-1/3 p-4">
          <h2 className="text-3xl font-bold mb-4">All Posts</h2>
          <ul className="space-y-4">
            {posts.map((post) => (
              <li key={post.id} className="bg-gray-100 p-4 rounded-lg">
                <p className="text-sm text-gray-600">By {post.author} | {post.date}</p>
                <h3 className="text-xl font-semibold">{post.title}</h3>
              </li>
            ))}
          </ul>
          <a href="#" className="text-blue-500 hover:underline mt-4 block text-center">View All Posts</a>
        </div>
      </div>
    </div>
  );
};

export default FeatureSection;
