import NavItem from "./NavItem";

function NavBar() {
  return (
    <div className="flex gap-3">
      <NavItem path="/">Home</NavItem>
      <NavItem path="https://github.com/darkphotonKN/age-of-carnath">
        Github
      </NavItem>
    </div>
  );
}

export default NavBar;
