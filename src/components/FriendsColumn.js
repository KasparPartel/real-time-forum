import { useEffect, useState } from "react";

const friend = class {
  constructor(name, status) {
    this.name = name;
    this.status = status;
  }
};

export default function FriendsColumn() {
  const [friends, setFriends] = useState([]);

  useEffect(() => {
    setFriends([
      {
        id: 1,
        name: "Kaspar Pärtel",
        online: true,
      },
      {
        id: 2,
        name: "Krister Riska",
        online: false,
      },
      {
        id: 3,
        name: "Romet Selgis",
        online: true,
      },
      {
        id: 4,
        name: "Henri Porila",
        online: true,
      },
      {
        id: 5,
        name: "Toomas Rüütel",
        online: false,
      },
      {
        id: 6,
        name: "Victoria Vunk",
        online: true,
      },
    ]);
  }, []);

  return (
    <div className="flex flex-col items-center sticky top-0 h-screen basis-52 py-6 bg-cyan-100">
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
