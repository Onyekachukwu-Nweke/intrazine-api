const FeaturedPost = ({ title, author, date, summary, imageUrl }) => {
    return (
        <div className="bg-white shadow-lg rounded-lg p-4 my-4">
            <img src={imageUrl} alt="post cover" className="rounded-t-lg w-full h-48 object-cover" />
            <div className="p-4">
                <h2 className="text-xl font-bold">{title}</h2>
                <p className="text-gray-600">By {author} | {date}</p>
                <p className="text-gray-800">{summary}</p>
                <button className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Read More
                </button>
            </div>
        </div>
    );
};

const PostCard = ({ title, author, date, summary, imageUrl }) => {
    return (
        <div className="bg-white shadow-md rounded-lg p-4">
            <img src={imageUrl} alt="post cover" className="rounded-t-lg w-full h-32 object-cover" />
            <div className="p-4">
                <h3 className="text-lg font-bold">{title}</h3>
                <p className="text-sm text-gray-600">By {author} | {date}</p>
                <p className="text-gray-800">{summary}</p>
                <button className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Read More
                </button>
            </div>
        </div>
    );
};

const BlogPage = () => {
    // Sample data with real images from Unsplash
    const posts = [{
        title: 'Step-by-step guide to choosing great font pairs',
        author: 'John Doe',
        date: 'May 23, 2022',
        summary: 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.',
        imageUrl: 'https://images.unsplash.com/photo-1561948955-570b270e7c36?fit=crop&w=600&q=80', // Example image from Unsplash
    }, {
        title: 'Design tips for designers that cover all the bases',
        author: 'Jane Smith',
        date: 'June 10, 2022',
        summary: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
        imageUrl: 'https://images.unsplash.com/photo-1556075798-4825dfaaf498?fit=crop&w=600&q=80', // Example image from Unsplash
    }];

    return (
        <div className="container mx-auto px-4">
            <FeaturedPost {...posts[0]} />
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {posts.map(post => <PostCard key={post.title} {...post} />)}
            </div>
        </div>
    );
};

export default BlogPage;