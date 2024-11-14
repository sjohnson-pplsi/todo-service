import { todoService } from "@/services/todo_service";
import { AppBar, Button, Container, Toolbar, Typography } from "@mui/material";
import { MessageForm } from "./message_form";

export default async function Home() {
  const todos = await todoService.listTodos(0, 0);

  return (
    <>
      <AppBar position="sticky">
        <Toolbar>Top</Toolbar>
      </AppBar>
      <Container>
        <Typography variant="h1">Title</Typography>
        <u>
          {todos.data.map((todo) => (
            <li key={todo.id}>{todo.note}</li>
          ))}
        </u>
        <MessageForm />
        <Button>Ok</Button>
      </Container>
    </>
  );
}
