import { useState } from 'react';
import { FaArrowLeft, FaArrowRight } from 'react-icons/fa';

const testimonials = [
  {
    text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
    author: "Jonathan Vallem",
    location: "New York, USA",
    image: "https://randomuser.me/api/portraits/men/32.jpg"
  },
  {
    text: "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
    author: "Jane Doe",
    location: "San Francisco, USA",
    image: "https://randomuser.me/api/portraits/women/32.jpg"
  },
  {
    text: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
    author: "John Smith",
    location: "Chicago, USA",
    image: "https://randomuser.me/api/portraits/men/33.jpg"
  }
];

const Testimonials = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  const prevTestimonial = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex === 0 ? testimonials.length - 1 : prevIndex - 1
    );
  };

  const nextTestimonial = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex === testimonials.length - 1 ? 0 : prevIndex + 1
    );
  };

  const { text, author, location, image } = testimonials[currentIndex];

  return (
    <div className="bg-light py-16 px-4">
      <div className="container bg-light-yellow mx-auto flex flex-col md:flex-row items-center justify-between p-8">
        {/* Left Section */}
        <div className="w-full md:w-1/3 mb-8 md:mb-0 text-center md:text-left">
          <h3 className="text-sm uppercase text-gray-600 mb-2">Testimonials</h3>
          <h2 className="text-2xl font-bold mb-4">What people say about our blog</h2>
          <p className="text-gray-700">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor.</p>
        </div>

        {/* Right Section */}
        <div className="w-full md:w-2/3 bg-white p-8 rounded-lg">
          <p className="text-lg text-left font-bold text-gray-800 mb-6 w-50">{text}</p>
          <div className="flex items-center">
            <img src={image} alt={author} className="w-12 h-12 rounded-full mr-4" />
            <div>
              <p className="text-lg font-semibold text-gray-900">{author}</p>
              <p className="text-sm text-gray-500">{location}</p>
            </div>
          </div>
          <div className="flex mt-6 justify-end space-x-4">
            <button
              onClick={prevTestimonial}
              className="p-2 rounded-full bg-gray-300 hover:bg-gray-400 focus:outline-none"
            >
              <FaArrowLeft />
            </button>
            <button
              onClick={nextTestimonial}
              className="p-2 rounded-full bg-black text-white hover:bg-gray-700 focus:outline-none"
            >
              <FaArrowRight />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Testimonials;
