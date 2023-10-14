import reactLogo from './assets/react.svg'
import telegramLogo from './assets/telegram.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const onShowTelegramDialog = () => {
    // Example of interaction with telegram-web-app.js script to show a native alert
    window.Telegram.WebApp.showAlert("Message from Telegram Mini App Template Bot");
  };

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank" rel="noreferrer">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank" rel="noreferrer">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
        <a href="https://core.telegram.org/bots/webapps" target="_blank" rel="noreferrer">
          <img src={telegramLogo} className="logo telegram" alt="Telegram logo" />
        </a>
      </div>
      <h1>Telegram Mini App Template Bot</h1>
      <div className="card">
        <button onClick={onShowTelegramDialog}>
          Show Telegram Dialog
        </button>
      </div>
      <p className="read-the-docs">
        Click on the logos to learn more
      </p>
    </>
  )
}

export default App
