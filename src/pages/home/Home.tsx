import { Link } from 'react-router-dom'
import './Home.css'

const Home = () => {
  return (
    <main className="page-main">
      <section id="home" className="home-section">
        <h2>Welcome</h2>
        <p>このサイトは現在開発中です。</p>
        <div className="home-actions">
          <Link to="/hue-are-you" className="home-action home-action--primary">
            Hue Are You? を試す
          </Link>
          <Link to="/portfolio" className="home-action home-action--secondary">
            ポートフォリオを見る
          </Link>
        </div>
      </section>
    </main>
  )
}

export default Home
