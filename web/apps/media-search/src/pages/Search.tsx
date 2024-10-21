// Copyright 2024 Google, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import SearchIcon from "@mui/icons-material/Search";
import { Grid2, IconButton, InputBase, Paper, Snackbar, SnackbarCloseReason, TextField, Typography } from "@mui/material";
import axios from "axios";
import { useState } from "react";

interface CastMember {
  actor_name: string;
  character_name: string;
}

interface Scene {
  sequence: number;
  start: string;
  end: string;
  script: string;
}

interface MediaResult {
  id: string;
  create_date: Date;
  title: string;
  summary: string;
  director: string;
  release_year: number;
  genre: string;
  rating: string;
  cast: CastMember[];
  scenes: Scene[];
}

const formatScript = (val: string): string => {
  return val.replace("\n", "<br/>")
}

const SceneData = ({ scene }: { scene: Scene }) => {
  return( 
  <>
    <Grid2 size={2}>{scene.sequence}</Grid2>
    <Grid2 size={5}>{scene.start}</Grid2>
    <Grid2 size={5}>{scene.end}</Grid2>
    <Grid2 size={12} sx={{textAlign: 'left'}}><div dangerouslySetInnerHTML={{ __html: formatScript(scene.script) }} /></Grid2>
  </>)
};

const MediaRow = ({ result }: { result: MediaResult }) => {
  return (
    <>
      <Grid2 size={4} sx={{textAlign: 'left'}}>
        <Typography variant="h6">{result.title}</Typography>
        <Typography variant="caption">{result.summary}</Typography>
      </Grid2>
      <Grid2 size={8}>
        {result.scenes.map((s: Scene) => (
          <Grid2 container spacing={2} sx={{mb: 2}}>
            <SceneData key={`${result.id}-${s.sequence}` } scene={s} />
          </Grid2>
        ))}
      </Grid2>
    </>
  );
};

const Search = () => {
  const [query, setQuery] = useState<string>(null!);
  const [results, setResults] = useState<Array<MediaResult>>([]);
  const [error, setError] = useState<string>(null!);
  const [open, setOpen] = useState<boolean>(false);

  const runQuery = () => {
    axios
      .get(`http://localhost:8080/api/v1/media?count=5&s=${query}`)
      .then((r) => {
        console.log(r)
        if (r.status == 200) {
            setResults([...r.data]);
        } else {
          setError(
            `Invalid HTTP Response: ${r.status} ${r.statusText} - ${r.data}`,
          );
          setOpen(true);
        }
      })
      .catch((e) => {
        setError(e);
        setOpen(true);
      });
  };

  const handleClose = (
    _: React.SyntheticEvent | Event,
    reason?: SnackbarCloseReason,
  ) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpen(false);
  };

  return (
    <>
      <Snackbar
        open={open}
        autoHideDuration={6000}
        onClose={handleClose}
        message={error}
        action={<></>}
      />
      <Paper sx={{p: 2, mb: 2}}>
        <TextField
          variant="outlined"
          sx={{ ml: 1, flex: 1 }}
          placeholder="Search Media"
          value={query}
          onChange={(v) => setQuery(v.target.value)}
        />
        <IconButton sx={{ p: "10px" }} aria-label="search" onClick={runQuery}>
          <SearchIcon />
        </IconButton>
      </Paper>

      {results && results.length > 0 ? (
        <Paper sx={{p: 2, mb: 2}}>
          <Grid2 container spacing={2}>
            {results.map((r) => (
              <MediaRow key={r.id} result={r} />
            ))}
          </Grid2>
        </Paper>
      ) : (
        <></>
      )}
    </>
  );
};

export default Search;
