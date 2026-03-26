import styles from "./Input.css";

const Input = ({ label, register, name, error, type = "text", ...rest }) => {
  return (
    <div className={styles.inputGroup}>
      {label && <label className={styles.label}>{label}</label>}

      <input
        className={styles.input}
        type={type}
        {...register(name, type === "number" ? { valueAsNumber: true } : {})}
        {...rest}
      />

      {error && <p className={styles.error}>{error.message}</p>}
    </div>
  );
};

export default Input;