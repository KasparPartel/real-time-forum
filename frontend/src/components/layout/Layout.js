import MainNavigation from './MainNavigation';
import PagesNavigation from './PagesNavigation';
import Userlist from './Userlist';
import classes from './Layout.module.css';

function Layout(props) {
    return (
        <div>
            <MainNavigation />
            <PagesNavigation loginState={props.loginState}/>
            <main className={classes.main}>{props.children}</main>
            <Userlist/>
        </div>
    );
}

export default Layout;
