import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { registerSchema } from "../../forms/registerSchema";

import Input from "../../components/ui/Input/Input";
import Button from "../../components/ui/Button/Button";

import "../../styles/Register.css";

const Register = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(registerSchema),
  });

  const onSubmit = (data) => {
    console.log("Form Data:", data);
  };

  return (
    <div className="pageWrapper">
      <form onSubmit={handleSubmit(onSubmit)} className="registercontainer">
        <h2>Register</h2>

        <Input
          label="Name" name="name" register={register} error={errors.name} placeholder="Enter your name"
 />

        <Input
          label="Email" name="email" register={register} error={errors.email}  placeholder="Enter your email"/>

        <Input
          label="Password" name="password" type="password" register={register} error={errors.password} placeholder="Enter password"/>

        <Input
          label="Age" name="age" type="number" register={register} error={errors.age} placeholder="Enter age"/>

        <Button text="Submit" type="submit" />
      </form>
    </div>
  );
};

export default Register;