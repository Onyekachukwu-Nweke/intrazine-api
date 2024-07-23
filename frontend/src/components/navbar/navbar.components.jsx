const Navbar = () => {
  return (
    <nav className='flex flex-col md:flex-row bg-black justify-between items-center p-4 text-light'>
      <div className='flex-shrink-0'>
        <a href="#" className='text-light italic text-2xl font-bold'>Piko Blog</a>
      </div>

      <div className='flex flex-col items-center md:flex-row md:ml-auto mt-4 md:mt-0'>
        <a href="#" className='text-light hover:text-gray-400 mx-2 md:mx-4'>Home</a>
        <a href="#" className='text-light hover:text-gray-400 mx-2 md:mx-4'>Blog</a>
        <a href="#" className='text-light hover:text-gray-400 mx-2 md:mx-4'>About Us</a>
        <a href="#" className='text-light hover:text-gray-400 mx-2 md:mx-4'>Contact us</a>
        <a href="#" className='text-light hover:text-gray-400 mx-2 md:mx-4'>
          <button className='bg-light text-black px-4 py-2 hover:bg-gray-200'>Subscribe</button>
        </a>
      </div>
    </nav>
  );
};

export default Navbar;
