import { Link } from 'react-router-dom';

import classes from './PagesNavigation.module.css';
// import FavoritesContext from '../../store/favorites-context';

function PagesNavigation() {
  
  return (
    
      
      <nav className={classes.pages}>
        <ul>
          <li>
            <Link to='/login'>Log in</Link>
          </li>
          <li>
            <Link to='/register'>Register</Link>
          </li>
          <li>
            <Link to='/'>Feed</Link>
          </li>
          <li>
            <Link to='/create-post'>Create new post</Link>
          </li>
          <li>
            <Link to='/messages'>Private messages</Link>
          </li>
          <li>
            <Link to='/profile'>User profile</Link>
          </li>
          
        </ul>
      </nav>

  );
}

export default PagesNavigation;