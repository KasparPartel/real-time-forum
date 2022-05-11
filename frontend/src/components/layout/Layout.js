import MainNavigation from './MainNavigation';
import PagesNavigation from './PagesNavigation';
import Userlist from './Userlist';

import classes from './Layout.module.css';

function Layout({user, children}) {
    return (
        <div>
            <MainNavigation userName={user} />
            <PagesNavigation user={user}/>
            <main className={classes.main}>{children}</main>
            <Userlist/>
        </div>
    );
}

export default Layout;
