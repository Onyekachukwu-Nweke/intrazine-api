import { ReactSVG } from 'react-svg';
import GroupSvg from '../../assets/yellow.svg'; // Ensure this path is correct

const AboutUsAndMission = () => {
  return (
    <div className="relative container mx-auto my-16 px-4">
      {/* SVG Decoration */}
      {/* <div className="absolute -translate-y-[1.5rem] translate-x-[27.5rem]">
        <ReactSVG src={GroupSvg} className="w-32 h-32 md:w-48 md:h-48" />
      </div> */}

      <div className="absolute hidden md:block -translate-y-[1.5rem] md:-translate-y-[1.5rem] lg:translate-x-[11.5rem] md:translate-x-[11.5rem] ">
        <ReactSVG src={GroupSvg} className="w-32 h-32 md:w-48 md:h-48" />
      </div>

      <div className="flex flex-col md:flex-row justify-between bg-lavender p-8 md:p-12">
        {/* About Us */}
        <div className="md:w-1/2 p-4 md:p-8">
          <h2 className="text-sm font-semibold uppercase mb-2 text-black">About Us</h2>
          <h3 className="text-2xl md:text-3xl font-bold font-weight-700 mb-4 tracking-tighter leading-snug md:leading-10">
            We are a community of content writers who share their learnings
          </h3>
          <p className="text-gray-700 mb-4">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </p>
          <a href="#" className="text-purple-600 hover:underline">Read More</a>
        </div>

        {/* Our Mission */}
        <div className="md:w-1/2 p-4 md:p-8 mt-8 md:mt-0">
          <h2 className="text-sm font-semibold uppercase mb-2 text-black">Our Mission</h2>
          <h3 className="text-2xl md:text-3xl font-bold font-weight-700 mb-4 tracking-tighter leading-snug md:leading-10">
            Creating valuable content for creatives all around the world
          </h3>
          <p className="text-gray-700 mb-4">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </p>
        </div>
      </div>
    </div>
  );
};

export default AboutUsAndMission;
