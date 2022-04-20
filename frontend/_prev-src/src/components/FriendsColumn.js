import { useEffect, useState } from "react";

// const friend = class {
//   constructor(name, status) {
//     this.name = name;
//     this.status = status;
//   }
// };

const randomNames = [
  "Adina Oates",
  "Moira Gooch",
  "Joann Godoy",
  "Brenna Ames",
  "Roman Provencher",
  "Harmony Francis",
  "Hayleigh Dias",
  "Tory Shropshire",
  "Hamzah Diep",
  "Brailynn Beeman",
  "Kierra Fontaine",
  "Sakura Corwin",
  "Jude Ahmad",
  "Philomena Loper",
  "Julietta Branson",
  "Emerie Boswell",
  "Sapphira Mcmullin",
  "Nikita Pitt",
  "Jadelyn Frizzell",
  "Amyrah Spinks",
];

export default function FriendsColumn() {
  const [friends, setFriends] = useState([]);

  useEffect(() => {
    // requestNames();
    setFriends(
      randomNames.map((name, index) => {
        if (Math.floor(Math.random() * 2)) {
          return { id: index, name, online: true };
        } else {
          return { id: index, name, online: false };
        }
      })
    );
  }, []);

  // function requestNames() {
  //   fetch("https://uinames.com/api/?amount=25")
  //     .then((response) => response.json())
  //     .then((json) => {
  //       setFriends(
  //         json.map((name) => {
  //           if (Math.floor(Math.random() * 2)) {
  //             return { name, online: true };
  //           } else {
  //             return { name, online: false };
  //           }
  //         })
  //       );
  //     });
  // }

  return (
    <div className="flex flex-col items-center shrink-0 sticky top-0 h-screen basis-52 py-6 bg-red-100">
      <div className="mb-10">
        <span className="text-green-500">&#x25CF; </span>
        <h3 className="inline underline font-bold">Online</h3>
        <ul>
          {friends.map((f) => {
            if (f.online) {
              return <li key={f.id}>{f.name}</li>;
            }
          })}
        </ul>
      </div>

      <div>
        <span className="text-red-500">&#x25CF; </span>
        <h3 className="inline underline font-bold">Offline</h3>
        <ul>
          {friends.map((f) => {
            if (!f.online) {
              return <li key={f.id}>{f.name}</li>;
            }
          })}
        </ul>
      </div>
    </div>
  );
}
