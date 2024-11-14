import { listTodos } from "@/services/todo_service";
import {
  AppBar,
  Button,
  Container,
  List,
  Toolbar,
  Typography,
} from "@mui/material";

import { TodoListItem } from "@/components/todo_list_item";

import { MessageForm } from "./message_form";

export default async function Home({
  params: { limit = "0", offset = "0" },
}: {
  params: { limit: string; offset: string };
}) {
  const todos = await listTodos(parseInt(limit), parseInt(offset));

  console.log(todos);

  return (
    <>
      <AppBar position="sticky">
        <Toolbar>Top</Toolbar>
      </AppBar>
      <Container>
        <Typography variant="h1">Title</Typography>
        <List>
          {todos.data.map((todo) => (
            <TodoListItem key={todo.id} todo={todo} />
          ))}
        </List>
        <MessageForm />
        <Button>Ok</Button>
      </Container>
    </>
  );
}
