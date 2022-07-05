import { useContext, useEffect, useState } from "react";
import { UserContext } from "../UserContext";
import { useNavigate } from "react-router-dom";

import styles from "./CreatePost.module.css";
import globalStyles from "../App.module.css";

export default function CreatePost() {
  const { user } = useContext(UserContext);
  const navigate = useNavigate();

  if (!user) {
    navigate("/login", { replace: true });
  }

  const [formData, setFormData] = useState({
    user_id: user.id.toString(),
    filename: "",
  });
  const [categories, setCategories] = useState([]);
  // const [selectedCategoryID, setSelectedCategoryID] = useState(null);

  useEffect(() => {
    getCategories();
  }, []);

  const getCategories = () => {
    fetch(`http://localhost:4000/v1/api/categories/`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setCategories(data);
        console.log("categories", categories);
      });
  };

  // handleChange sets new object based on changed form data
  const handleChange = (e) => {
    let formDataCopy = Object.assign({}, formData);
    let name = e.target.getAttribute("name");

    formDataCopy[name] = e.target.value;
    setFormData(formDataCopy);
  };

  // handleSubmit sends data to api and navigates to new post page
  const handleSubmit = (e) => {
    e.preventDefault();

    console.log(formData);

    fetch("http://localhost:4000/v1/api/post/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else {
          console.log("Cannot add post!");
          return;
        }
      })
      .then((data) => navigate(`/post/${data}`, { replace: true }));
  };

  return (
    <div className={`${globalStyles.flex} ${styles.container} `}>
      <h2 className={styles.header}>Create your new post here</h2>

      <form className={styles.form} onSubmit={handleSubmit}>
        <label className={styles.form__label} htmlFor="title">
          Post title
        </label>
        <br />
        <input
          type="text"
          className={styles.form__input}
          name="title"
          onChange={handleChange}
          required
        />
        <br />
        <select
          name="category_id"
          className={styles.form__select}
          defaultValue=""
          onChange={handleChange}
          required
        >
          <option value="" disabled>
            Choose Category
          </option>
          {categories?.map((category) => {
            return (
              <option value={category.id} key={category.id}>
                {category.title}
              </option>
            );
          })}
        </select>
        <br></br>
        <textarea
          className={styles.form__textarea}
          name="body"
          rows="4"
          cols="70"
          placeholder="Enter your post here"
          onChange={handleChange}
          required
        ></textarea>
        <br></br>

        <input
          type="submit"
          className={styles.btn__submit}
          value="Create Post"
        />
      </form>
    </div>
  );
}
