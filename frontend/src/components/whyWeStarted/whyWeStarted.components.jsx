import startedImage from '../../assets/started.png';

const WhyWeStarted = () => {
  return (
    <div className="relative container mx-auto my-16 px-4">
      <div className="flex flex-col md:flex-row items-center">
        {/* Image Section */}
        <div className="md:w-2/3 relative">
          <img
            src={startedImage} // Example image from Unsplash
            alt="Group of people"
            className="w-full h-full object-cover"
          />
          {/* Text Section (Overlay) */}
          <div className="absolute right-0 top-0 md:translate-x-1/2 md:translate-y-1/2 bg-light p-12 md:w-2/3">
            <h3 className="text-sm font-semibold uppercase text-gray-500 mb-2">Why We Started</h3>
            <h2 className="text-3xl md:text-4xl font-bold text-gray-800 mb-4">
              It started out as a simple idea and evolved into our passion
            </h2>
            <p className="text-gray-700 mb-6">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
            </p>
            <a href="#" className="inline-block bg-yellow text-gray-800 font-bold py-2 px-4 hover:bg-yellow-300">
              Discover our story
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default WhyWeStarted;