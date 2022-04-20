import classes from './Userlist.module.css';

function Userlist() {
  
  return (
    <ul className={classes.userlist}>
        <li className={classes.online}>HC: Online user</li>
        <li className={classes.offline}>HC: Offline user</li>
    </ul>
  );
}

export default Userlist;