import React from 'react';
import logo from 'assets/images/logo.svg';
import styles from './App.module.scss';
import { addMessage } from 'app/actions/messages';
import Messages from "app/layout/Messages";
import { useDispatch } from 'react-redux';
import { ADD_MESSAGE } from 'app/types/messages';

const App: React.FC = () => {

  const dispatch = useDispatch();

  return (
    <div className={styles["App"]}>
      <header className={styles["App-header"]}>
        <img src={logo} className={styles["App-logo"]} alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className={styles["App-link"]}
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <Messages/>
        <button onClick={() => dispatch({type:ADD_MESSAGE, payload: {text:"test", timeout:5 } })}>ADD MESSAGE TEST</button>
      </header>
    </div>
  );
}

export default App;