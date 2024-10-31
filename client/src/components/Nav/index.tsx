import NavItem from "./NavItem";

function NavBar() {
  return (
    <div className="flex justify-between">
      <div className="flex gap-3">
        <NavItem path="/">Home</NavItem>
        <NavItem path="https://github.com/darkphotonKN/age-of-carnath">
          Github
        </NavItem>
      </div>
      <div className="flex gap-3">
        <NavItem path="/account">Account</NavItem>
        <NavItem path="/signin">Sign In</NavItem>
      </div>
    </div>
  );
}

export default NavBar;
