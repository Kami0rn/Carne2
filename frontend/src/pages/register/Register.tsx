import React, { useState, ChangeEvent, FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import { UserInterface } from "../../interfaces/Iuser";
import { CreateUser } from "../../services/http/user/user";

function Register() {
  const [inputs, setInputs] = useState<UserInterface>({
    PhoneNumber: 0, // or any default number value
  });
  
  const navigate = useNavigate();
  const MySwal = withReactContent(Swal);

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setInputs((values) => ({
      ...values,
      [name]: name === "PhoneNumber" ? parseInt(value, 10) || 0 : value,
    }));
  };
  
  

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    console.log(inputs);

    try {
      const result = await CreateUser({
        FirstName: inputs.FirstName,
        LastName: inputs.LastName,
        PhoneNumber: inputs.PhoneNumber,
        Email: inputs.Email,
        Password: inputs.Password,
        Wallet: inputs.Wallet,
      });

      if (result && result.status === "ok") {
        MySwal.fire({
          html: <i>{result.message}</i>,
          icon: "success",
        }).then((value) => {
          navigate("/login");
        });
      } else {
        const message =
          result && result.message ? result.message : "No result message";
        MySwal.fire({
          html: <i>{message}</i>,
          icon: "error",
        });

        console.log(result);

      }
    } catch (error) {
      MySwal.fire({
        html: <i>{String(error)}</i>,
        icon: "error",
      });
    }
  };

  return (
    <div className="w-full max-w-xs">
      
      <form onSubmit={handleSubmit} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <label>
        FirstName:
          <input
            type="text"
            name="FirstName"
            value={inputs.FirstName}
            onChange={handleChange}
          />
        </label>
        <label>
        LastName:
          <input
            type="text"
            name="LastName"
            value={inputs.LastName}
            onChange={handleChange}
          />
        </label>
        <label>
        PhoneNumber:
          <input
            type="number"
            name="PhoneNumber"
            value={inputs.PhoneNumber}
            onChange={handleChange}
          />
        </label>
        <label>
        Email:
          <input
            type="text"
            name="Email"
            value={inputs.Email }
            onChange={handleChange}
          />
        </label>
        <label>
        Password:
          <input
            type="password"
            name="Password"
            value={inputs.Password }
            onChange={handleChange}
            required // Add the 'required' attribute
          />
        </label>

        <input type="submit" />
      </form>
    </div>
  );
}

export default Register;
