import React from "react";

const User = () => {
  const [name, setName] = React.useState("");
  const onChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };
  const onClick = async () => {
    const res = await fetch("https://api.test.local:8081/", {
      method: "GET",
      headers: {
        "content-type": "application/json",
      },
      mode: "cors",
      credentials: "include",
    });
  };

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await fetch("https://api.test.local:8081/name", {
      method: "POST",
      headers: {
        "content-type": "application/json",
      },
      mode: "cors",
      credentials: "include",
      body: JSON.stringify({ name }),
    });
  };

  return (
    <>
      <button onClick={onClick}>set cookie</button>
      <form onSubmit={onSubmit}>
        <label>type your name </label>
        <input type="text" value={name} onChange={onChangeName} />
        <button type="submit">send</button>
      </form>
    </>
  );
};

export default User;
