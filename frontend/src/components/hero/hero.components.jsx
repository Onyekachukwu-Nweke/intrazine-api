import heroImage from '../../assets/hero.png'; // Ensure this path is correct

const HeroSection = () => {
  return (
    <div className="relative h-screen bg-cover bg-center" style={{ backgroundImage: `url(${heroImage})` }}>
      <div 
        className="absolute inset-0 flex items-center"
        style={{ background: 'radial-gradient(80.99% 71.93% at 74.58% 0%, rgba(0, 0, 0, 0.00) 0%, rgba(0, 0, 0, 0.60) 100%)' }}
      >
        <div className="container mx-auto px-4">
          <div className="max-w-lg text-left text-light">
            <h1 className="text-4xl font-bold mb-4">Welcome to Piko Blog</h1>
            <p className="text-lg mb-6">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus imperdiet, nulla et dictum interdum, nisi lorem egestas odio, vitae scelerisque enim ligula venenatis dolor.
            </p>
            <button className="bg-yellow text-black px-6 py-3 rounded hover:bg-gray-200">Read More</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HeroSection;
