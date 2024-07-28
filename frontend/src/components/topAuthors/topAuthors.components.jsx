import { FaFacebook, FaXTwitter, FaInstagram, FaLinkedin, FaHeart, FaEye } from "react-icons/fa6";

const authors = [
  {
    id: 1,
    name: 'Floyd Miles',
    image: 'https://randomuser.me/api/portraits/men/1.jpg',
    reads: 1200,
    likes: 340,
    social: {
      facebook: '#',
      twitter: '#',
      instagram: '#',
      linkedin: '#',
    },
  },
  {
    id: 2,
    name: 'Dianne Russell',
    image: 'https://randomuser.me/api/portraits/women/29.jpg',
    reads: 1500,
    likes: 450,
    social: {
      facebook: '#',
      twitter: '#',
      instagram: '#',
      linkedin: '#',
    },
  },
  {
    id: 3,
    name: 'Jenny Wilson',
    image: 'https://randomuser.me/api/portraits/women/2.jpg',
    reads: 1100,
    likes: 300,
    social: {
      facebook: '#',
      twitter: '#',
      instagram: '#',
      linkedin: '#',
    },
  },
  {
    id: 4,
    name: 'Leslie Alexander',
    image: 'https://randomuser.me/api/portraits/men/5.jpg',
    reads: 900,
    likes: 250,
    social: {
      facebook: '#',
      twitter: '#',
      instagram: '#',
      linkedin: '#',
    },
  },
];

const TopAuthors = () => {
  return (
    <div className="container mx-auto my-16 px-4 text-center">
      <h2 className="text-3xl font-bold mb-8">Top Authors of the Month</h2>
      <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
        {authors.map((author) => (
          <div key={author.id} className="bg-gray-100 p-6 rounded-lg hover:bg-light-yellow transition duration-300 ease-in-out">
            <img
              src={author.image}
              alt={author.name}
              className="w-24 h-24 rounded-full mx-auto mb-4"
            />
            <h3 className="text-xl font-bold mb-2">{author.name}</h3>
            <div className="flex justify-center items-center text-gray-700 mb-4">
              <FaEye className="mr-1" /> {author.reads} <span className="mx-2">|</span> <FaHeart className="mr-1" /> {author.likes}
            </div>
            <div className="flex justify-center space-x-4 text-gray-600">
              <a href={author.social.facebook} className="text-black"><FaFacebook /></a>
              <a href={author.social.twitter} className="text-black"><FaXTwitter /></a>
              <a href={author.social.instagram} className="text-black"><FaInstagram /></a>
              <a href={author.social.linkedin} className="text-black"><FaLinkedin /></a>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default TopAuthors;
