// axios.get("http://localhost:8080/materialList").then(function (response) {
//   console.log(response);
// });

let formData = new FormData();
let materialName = document.querySelector("#materialName");
let imageFile = document.querySelector("#file");
let lastFolderCreated;

document.querySelector(".b1").addEventListener("click", function (e) {
  e.preventDefault();

  if (materialName.value === undefined || materialName.value === "") {
    console.log("Error!!! Form fields are empty");
  } else {
    lastFolderCreated = materialName.value.replace(/\s/g, "");
    console.log(lastFolderCreated);
    formData.append("materialName", lastFolderCreated);
    console.log(formData);

    localStorage.setItem("lastFolderCreated", lastFolderCreated);

    axios
      .post("http://localhost:8080/newMaterial", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        // handle error
        console.log(error);
      });
  }
});

document.querySelector(".b2").addEventListener("click", function (e) {
  e.preventDefault();
  let sizeOfThePic = document.querySelector("select").value;
  lastFolderCreated = localStorage.getItem("lastFolderCreated");

  if (
    imageFile.files === undefined ||
    lastFolderCreated === undefined ||
    lastFolderCreated === "" ||
    sizeOfThePic === undefined ||
    sizeOfThePic === ""
  ) {
    console.log("Error!!! Form fields are empty");
  } else {
    // formData.delete(lastFolderCreated);
    console.log(imageFile.files[0]);
    formData.append("imageFile", imageFile.files[0]);
    console.log(lastFolderCreated);

    axios
      .post(
        "http://localhost:8080/newImage/" +
          lastFolderCreated +
          "/" +
          sizeOfThePic,
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      )
      .catch(function (error) {
        // handle error
        console.log(error);
      });
  }
});

// let testUrl = "test/medium.png";
// let formDatTest = new FormData();
// formDatTest.append("imageLoc", testUrl);

// axios
//   .post("http://localhost:8080/delete", formDatTest, {
//     headers: {
//       "Content-Type": "multipart/form-data",
//     },
//   })
//   .catch(function (error) {
//     // handle error
//     console.log(error);
//   });
