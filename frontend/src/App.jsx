import Navbar from './components/navbar/navbar.components';
import HeroSection from './components/hero/hero.components';
import FeatureSection from './components/feature/feature.component';
import AboutUsAndMission from './components/aboutAndMission/aboutUs.component';
import CategorySection from './components/categorySection/categories.component';
import WhyWeStarted from './components/whyWeStarted/whyWeStarted.components';
import TopAuthors from './components/topAuthors/topAuthors.components';
import FeaturedIn from './components/clout/clout.components';

function App() {

  return (
    <>
      <Navbar />
      <HeroSection />
      <FeatureSection />
      <AboutUsAndMission />
      <CategorySection />
      <WhyWeStarted />
      <TopAuthors />
      <FeaturedIn />
    </>
  )
}

export default App
