import {useContext} from "react";
import {UserContext} from "../UserContext";
import {useNavigate} from "react-router-dom";


function Messages() {
  const {user} = useContext(UserContext)
  const navigate = useNavigate()

  if (!user) {
    navigate("/login", {replace: true})
  }
 
  return (
    <div>Messages div</div> 
  );
}
  
export default Messages;