import { ReactSVG } from 'react-svg';
import Logo1 from '../../assets/Logo1.svg';
import Logo2 from '../../assets/Logo2.svg';
import Logo3 from '../../assets/Logo3.svg';
import Logo4 from '../../assets/Logo4.svg';
import Logo5 from '../../assets/Logo5.svg';

const FeaturedIn = () => {
  return (
    <div className="container mx-auto my-16 px-4">
      <div className="flex flex-col md:flex-row items-center justify-center md:justify-between space-y-4 md:space-y-0 md:space-x-8">
        <div className="text-center md:text-left md:flex-shrink-0">
          <p className="text-gray-500 text-sm">We are</p>
          <h2 className="text-lg font-semibold text-gray-700">Featured in</h2>
        </div>
        <div className="flex flex-wrap justify-center md:justify-start items-center space-x-4 md:space-x-8">
          <ReactSVG src={Logo1} className="m-2 text-gray-500" />
          <ReactSVG src={Logo2} className="m-2 text-gray-500" />
          <ReactSVG src={Logo3} className="m-2 text-gray-500" />
          <ReactSVG src={Logo4} className="m-2 text-gray-500" />
          <ReactSVG src={Logo5} className="m-2 text-gray-500" />
        </div>
      </div>
    </div>
  );
};

export default FeaturedIn;
