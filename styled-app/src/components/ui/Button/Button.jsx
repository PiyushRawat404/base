import styles from "./Button.css";

const Button = ({ text, type = "button" }) => {
  return (
    <button className={styles.button} type={type}>
      {text}
    </button>
  );
};

export default Button;