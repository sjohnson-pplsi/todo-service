import { todoService } from "@/services/todo_service";
import { AppBar, Button, Container, Toolbar, Typography } from "@mui/material";

export default async function Home() {
  const message = await todoService.sayHello("name");
  return (
    <>
      <AppBar position="sticky">
        <Toolbar>Top</Toolbar>
      </AppBar>
      <Container>
        <Typography variant="h1">Title</Typography>
        <p>Message: {message}</p>
        <Button>Ok</Button>
      </Container>
    </>
  );
}
