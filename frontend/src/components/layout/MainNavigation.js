import {Link} from 'react-router-dom';

import classes from './MainNavigation.module.css';

function MainNavigation({user}) {
    return (
        <header className={classes.header}>
            <div className={classes.logo}>real-time-forum</div>
            <nav>
                <ul>
                    <li>
                        <Link to='/'>All posts</Link>
                    </li>
                    <li>
                        <Link to='/create-post'>Add new post</Link>
                    </li>
                    <li>
                        <Link to="/profile">{user ? user.username : ""}</Link>
                    </li>
                </ul>
            </nav>
        </header>
    );
}

export default MainNavigation;
