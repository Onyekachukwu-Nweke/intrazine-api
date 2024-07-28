import { FaRegBuilding, FaRegNewspaper, FaRegStar, FaWater, FaCameraRetro } from 'react-icons/fa';

const FeaturedIn = () => {
  return (
    <div className="container mx-auto my-16 px-4 text-center">
      <div className="flex items-center justify-center space-x-8">
        <span className="text-gray-500 text-sm">We are</span>
        <h2 className="text-lg font-semibold text-gray-700">Featured in</h2>
        <div className="flex space-x-8">
          <FaRegBuilding className="text-gray-500 h-8 w-8" />
          <FaRegNewspaper className="text-gray-500 h-8 w-8" />
          <FaRegStar className="text-gray-500 h-8 w-8" />
          <FaWater className="text-gray-500 h-8 w-8" />
          <FaCameraRetro className="text-gray-500 h-8 w-8" />
        </div>
      </div>
    </div>
  );
};

export default FeaturedIn;
