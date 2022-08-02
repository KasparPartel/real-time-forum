import React, { useEffect, useState } from "react";
import FeedPost from "../components/layout/FeedPost";

import styles from "./Feed.module.css";

export default function Feed() {
  const [posts, setPosts] = useState([]);
  const [categories, setCategories] = useState([]);
  const [selectedCategoryID, setSelectedCategoryID] = useState(null);
  const [selectedPosts, setSelectedPosts] = useState([]);

  useEffect(() => {
    getPosts();
    getCategories();
  }, []);

  useEffect(() => {
    setSelectedPosts(posts);
  }, [posts]);

  useEffect(() => {
    filterPosts();
    console.log("selected posts", selectedPosts);
  }, [selectedCategoryID]);

  // getPosts fetches all posts from api
  const getPosts = () => {
    fetch("http://localhost:4000/v1/api/post/", {
      method: "GET",
    })
      .then((response) => {
        if (response.ok) {
          console.log("okay");
          return response.json();
        }
        throw new Error("Something went wrong");
      })
      .then((json) => setPosts(json))
      .catch((error) => console.log(error));
  };

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

  const filterPosts = () => {
    console.log("category id selected", selectedCategoryID);
    if (selectedCategoryID) {
      setSelectedPosts(
        posts.filter((post) => selectedCategoryID === post["category_id"])
      );
    } else {
      setSelectedPosts(posts);
    }
  };

  const handleSelectChange = (e) => {
    const value = e.target.value;

    value === "" ? setSelectedCategoryID(null) : setSelectedCategoryID(value);
    console.log("category id", selectedCategoryID);
  };

  return (
    <div>
      <label className={styles.dropdown__label} htmlFor="category_id">
        Category:{" "}
      </label>
      <select
        name="category_id"
        className={styles.dropdown}
        value={selectedCategoryID || ""}
        onChange={handleSelectChange}
      >
        <option value="">All</option>
        {categories?.map((category) => (
          <option value={category.id} key={category.id}>
            {category.title}
          </option>
        ))}
      </select>
      {selectedPosts?.map((post) => (
        <FeedPost key={post.id} json={post} />
      ))}
    </div>
  );
}
