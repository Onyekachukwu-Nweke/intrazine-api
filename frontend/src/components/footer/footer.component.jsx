import { FaFacebook, FaTwitter, FaInstagram, FaLinkedin } from 'react-icons/fa';

const Footer = () => {
  return (
    <footer className="bg-gray-900 text-gray-300 py-12 px-4">
      <div className="container mx-auto grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Logo and Navigation */}
        <div className="flex flex-col items-start">
          <h3 className="text-2xl font-bold mb-4 text-white">Finsweet</h3>
          <ul className="space-y-2">
            <li><a href="#" className="text-gray-400 hover:text-white">Home</a></li>
            <li><a href="#" className="text-gray-400 hover:text-white">Blog</a></li>
            <li><a href="#" className="text-gray-400 hover:text-white">About us</a></li>
            <li><a href="#" className="text-gray-400 hover:text-white">Contact us</a></li>
            <li><a href="#" className="text-gray-400 hover:text-white">Privacy Policy</a></li>
          </ul>
        </div>

        {/* Newsletter Subscription */}
        <div className="bg-gray-800 p-6 rounded-lg text-center md:text-left">
          <h4 className="text-lg font-semibold mb-4 text-white">Subscribe to our newsletter to get latest updates and news</h4>
          <div className="flex items-center justify-center md:justify-start">
            <input 
              type="email" 
              placeholder="Enter Your Email" 
              className="p-2 w-full md:w-auto flex-grow md:flex-grow-0 rounded-l-md focus:outline-none bg-gray-700 text-gray-300 placeholder-gray-400" 
            />
            <button className="bg-yellow text-gray-900 p-2 hover:bg-yellow transition duration-300">Subscribe</button>
          </div>
        </div>

        {/* Contact Information */}
        <div className="flex flex-col items-start space-y-4">
          <p>Finstreet 118 2561 Fintown</p>
          <p>Hello@finsweet.com 020 7993 2905</p>
          <div className="flex space-x-4 text-gray-400">
            <a href="#" className="hover:text-white"><FaFacebook /></a>
            <a href="#" className="hover:text-white"><FaTwitter /></a>
            <a href="#" className="hover:text-white"><FaInstagram /></a>
            <a href="#" className="hover:text-white"><FaLinkedin /></a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
