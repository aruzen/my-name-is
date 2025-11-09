import type { CSSProperties } from 'react'
import { Link } from 'react-router-dom'

const buttonStyle: CSSProperties = {
  display: 'inline-flex',
  alignItems: 'center',
  justifyContent: 'center',
  padding: 'clamp(0.85rem, 3vw, 1rem) clamp(1.4rem, 4vw, 2rem)',
  fontSize: 'clamp(1rem, 4vw, 1.2rem)',
  border: 'none',
  borderRadius: '10px',
  color: 'white',
  cursor: 'pointer',
  textDecoration: 'none',
  minWidth: 'min(260px, 100%)',
  flex: '1 1 240px',
  transition: 'transform 0.2s ease, box-shadow 0.2s ease'
}

const buttonGroupStyle: CSSProperties = {
  marginTop: '2rem',
  width: '100%',
  display: 'flex',
  gap: '1rem',
  flexWrap: 'wrap',
  justifyContent: 'center'
}

const Home = () => {
  return (
    <main>
      <section id="home">
        <h2>Welcome</h2>
        <p>このサイトは現在開発中です。</p>
        <div style={buttonGroupStyle}>
          <Link
            to="/hue-are-you"
            style={{
              ...buttonStyle,
              background: 'linear-gradient(45deg, #4ecdc4, #45b7d1)'
            }}
          >
            Hue Are You? を試す
          </Link>
          <Link
            to="/portfolio"
            style={{
              ...buttonStyle,
              background: 'linear-gradient(45deg, #e74c3c, #c0392b)'
            }}
          >
            ポートフォリオを見る
          </Link>
        </div>
      </section>
    </main>
  )
}

export default Home
