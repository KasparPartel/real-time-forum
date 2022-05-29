import MainNavigation from './MainNavigation';
import PagesNavigation from './PagesNavigation';
import Userlist from './Userlist';

import classes from './Layout.module.css';
import {useContext} from "react";
import {UserContext} from "../../UserContext";

function Layout({children}) {
    const {user} = useContext(UserContext)

    return (
        <div>
            <MainNavigation user={user}/>
            <PagesNavigation user={user}/>
            <main className={classes.main}>{children}</main>
            <Userlist user={user}/>
        </div>
    );
}

export default Layout;
