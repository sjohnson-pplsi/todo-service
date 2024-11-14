import { listTodos } from "@/services/todo_service";
import { AppBar, Container, Toolbar } from "@mui/material";

import { TodoList } from "@/components/todo_list";

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
        <Toolbar>Todo List</Toolbar>
      </AppBar>
      <Container disableGutters>
        <TodoList todos={todos.data} />
      </Container>
    </>
  );
}
