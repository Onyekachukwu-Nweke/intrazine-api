import { BrowserRouter, Routes, Route } from 'react-router-dom';

import Navbar from './components/navbar/navbar.components';
import Footer from './components/footer/footer.component';

import Home from './routes/home/home.routes';
import Blogs from './routes/blogs/blogs.routes';

function App() {

  return (
    <BrowserRouter>
      <Navbar />
        <Routes>
          <Route path='/' index element={<Home />} />
          {/* <Route path='/blogs' element={<Blogs />} /> */}
        </Routes>
      <Footer />
    </BrowserRouter>
  )
}

export default App

    // <>
    //   <Navbar />
    //   <HeroSection />
    //   <FeatureSection />
    //   <AboutUsAndMission />
    //   <CategorySection />
    //   <WhyWeStarted />
    //   <TopAuthors />
    //   <FeaturedIn />
    //   <Testimonials />
    //   <JoinUs />
    //   <Footer />
    // </>