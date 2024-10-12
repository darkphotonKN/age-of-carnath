import Link from "next/link";

type NavItemProps = {
  path: string;
  children: React.ReactNode;
};

function NavItem({ path, children }: NavItemProps) {
  return (
    <Link href={path}>
      <div className="border-b-[1px] border-black hover:text-accent transition-color duration-200 ease-in">
        {children}
      </div>
    </Link>
  );
}

export default NavItem;
