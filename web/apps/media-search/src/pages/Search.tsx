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


import { Snackbar, SnackbarCloseReason, Typography } from "@mui/material";
import { useState } from "react";
import { MediaResult } from "../shared/model";
import SearchBar from "../components/SearchBar";
import MediaResults from "../components/MediaResults";

const Search = () => {
  const [results, setResults] = useState<Array<MediaResult>>([]);
  const [error, setError] = useState<string>(null!);
  const [open, setOpen] = useState<boolean>(false);

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
      <SearchBar setError={setError} setOpen={setOpen} setResults={setResults} />
      <MediaResults results={results} />
      <Snackbar
        open={open}
        autoHideDuration={6000}
        onClose={handleClose}
        message={error}
        action={<></>}
      />
    </>
  );
};

export default Search;
