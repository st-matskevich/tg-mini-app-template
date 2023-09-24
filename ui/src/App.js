import logo from './logo.svg';
import './App.css';

function App() {
  const onShowTelegramDialog = () => {
    // Example of interaction with telegram-web-app.js script to show a native alert
    window.Telegram.WebApp.showAlert("Message from Telegram Mini App Template Bot");
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Welcome to the Telegram Mini App Template Bot</p>
        <div className="button" onClick={onShowTelegramDialog}>Show Telegram Dialog</div>
      </header>
    </div>
  );
}

export default App;
