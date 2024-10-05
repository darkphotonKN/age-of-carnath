import { Button } from "@/components/Button";
import Link from "next/link";

export default function Home() {
  return (
    <div className="flex flex-col justify-center content-center h-full">
      <Link href="/game">
        <Button variant="default" size="default">
          Start
        </Button>
      </Link>
    </div>
  );
}
