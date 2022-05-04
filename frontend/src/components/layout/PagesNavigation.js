import {Link} from 'react-router-dom';
import {Fragment} from "react";

import classes from './PagesNavigation.module.css';

function PagesNavigation({loginState}) {
    return (
        <nav className={classes.pages}>
            <ul>
                {loginState ? (
                    <li>
                        <Link to='/logout'>Logout</Link>
                    </li>
                ) : (
                    <Fragment>
                        <li>
                            <Link to='/login'>Log in</Link>
                        </li>
                        <li>
                            <Link to='/register'>Register</Link>
                        </li>
                    </Fragment>
                )}
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