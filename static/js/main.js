  // Prevent form from reloading page
  var form = document.getElementById("inputForm");
  function handleForm(event) {
    event.preventDefault();
  }
  form.addEventListener("submit", handleForm);

  //Hide/show alert
  var alertFName = document.getElementById("alert-fname");
  alertFName.style.display = "none";
  var alertLName = document.getElementById("alert-lname");
  alertLName.style.display = "none";
  var alertDob = document.getElementById("alert-dob");
  alertDob.style.display = "none";

  //Fun Facts
  function getFunFactsSection() {
    var dob = document.getElementById("dob").value;
    var firstName = document.getElementById("firstName").value;
    var lastName = document.getElementById("lastName").value;

    var today = new Date();
    var dd = String(today.getDate()).padStart(2, "0");
    var mm = String(today.getMonth() + 1).padStart(2, "0");
    var yyyy = today.getFullYear();
    today = yyyy + "-" + mm + "-" + dd;

    if (dob >= today) {
      alertDob.style.display = "block";
    } else if (firstName.trim() === "" || firstName.match(/\d+/g)!= null) {
      alertFName.style.display = "block";
    } else if (lastName.trim() === "" || lastName.match(/\d+/g)!= null) {
      alertLName.style.display = "block";
    } else if (form.checkValidity() == false) {
      var list = form.querySelectorAll(":invalid");
      for (var item of list) {
        item.focus();
        alertDob.style.display = "block";
      }
    } else {
      alertFName.style.display = "none";
      alertLName.style.display = "none";
      alertDob.style.display = "none";
      document.getElementById("funfact-results").style.display = "block";
      document.getElementById("zodiac-results").style.display = "none";
      document.getElementById("numerology-results").style.display = "none";
      fetch(
        "/facts?dob=" +
          dob +
          "&firstName=" +
          firstName +
          "&lastName=" +
          lastName
      )
        .then((response) => response.json())
        .then((data) => {
          document.getElementById("first_name").innerHTML = firstName;
          document.getElementById("date_of_birth").innerHTML = data.date_of_birth;

          if (data.generation != "Not found") {
            document.getElementById("generation").style.display = "inline";
            document.getElementById("generation").innerHTML = "<br/>That means that your generation is called " + data.generation + ".";
          }else{
            document.getElementById("generation").style.display = "none";
          }

          const countUp = new CountUp('year_old',0,data.year_old);
          countUp.start();

          if (data.year_old >1) {
              document.getElementById("year").innerHTML  = " years";
          }else{
            document.getElementById("year").innerHTML  = " year";
          }

          const countUpMonthOld  = new CountUp('month_old',0,data.month_old);
          countUpMonthOld.start();

          if (data.month_old >1) {
              document.getElementById("month").innerHTML =  " months";
          }else{
            document.getElementById("month").innerHTML = " month";
          }

          const countUpDayOld = new CountUp('day_old',0,data.day_old);
          countUpDayOld.start();
          
          if (data.day_old >1) {
            document.getElementById("day").innerHTML = " days";
          }else{
            document.getElementById("day").innerHTML = " day";
          }

          const countDayOnEarth = new CountUp('day_on_earth',0,data.number_of_days);
          countDayOnEarth.start();
          if (data.number_of_days > 1) {
            document.getElementById("day_earth").innerHTML = " days";
          }else{
            document.getElementById("day_earth").innerHTML = " day";
          }

          if (data.next_birthdate == 0) {
            document.getElementById("birthday_pre").style.display = "none";
            document.getElementById("birthday_aft").style.display = "none";
            document.getElementById("birthday_next").innerHTML = " today is your birthday? Happy birthday!!! ";
          }else if (data.next_birthdate == 1) {
            document.getElementById("birthday_pre").style.display = "none";
            document.getElementById("birthday_aft").style.display = "none";
            document.getElementById("birthday_next").innerHTML = " there is " + data.next_birthdate+" day until your " + data.number_next_birthdate + " birthday?";
          }else{
            const countUpBirthDay = new CountUp('birthday_next',0,data.next_birthdate);
            countUpBirthDay.start();
            document.getElementById("birthday_pre").style.display = "inline";
            document.getElementById("birthday_aft").style.display = "inline";
            document.getElementById("birthday_pre").innerHTML = " there are ";
            document.getElementById("birthday_aft").innerHTML = " days until your " + data.number_next_birthdate + " birthday?";
          }
          document.getElementById("birthday_next_repeat").innerHTML = data.number_next_birthdate;
          document.getElementById("day_of_week_birthdate").innerHTML = data.day_of_week_birthdate;
          document.getElementById("day_of_week_next_birthdate").innerHTML = data.day_of_week_next_birthdate;
          const countUpCandles = new CountUp('number_of_candles',0,data.number_of_candles);
          countUpCandles.start();
          document.getElementById("date_fact").innerHTML = "On your birthdate in " + data.date_fact.year + ", " + data.date_fact.text + ".";
          
          location.hash = "#funfact-results";

            //function that changes the hash
            function setHash(newHash) {
                location.hash = 'someHashThatDoesntExist';
                location.hash = newHash;
            }
            setHash('#funfact-results');
        });
    }
  }

  //Zodiac
  function getZodiacSection() {
    var dob = document.getElementById("dob").value;
    var firstName = document.getElementById("firstName").value;
    var lastName = document.getElementById("lastName").value;

    var today = new Date();
    var dd = String(today.getDate()).padStart(2, "0");
    var mm = String(today.getMonth() + 1).padStart(2, "0");
    var yyyy = today.getFullYear();
    today = yyyy + "-" + mm + "-" + dd;

    if (dob >= today) {
      alertDob.style.display = "block";
    } else if (firstName.trim() === "" || firstName.match(/\d+/g)!= null) {
      alertFName.style.display = "block";
    } else if (lastName.trim() === "" || lastName.match(/\d+/g)!= null) {
      alertLName.style.display = "block";
    } else if (form.checkValidity() == false) {
      var list = form.querySelectorAll(":invalid");
      for (var item of list) {
        item.focus();
        alertDob.style.display = "block";
      }
    } else {
      alertFName.style.display = "none";
      alertLName.style.display = "none";
      alertDob.style.display = "none";
      document.getElementById("funfact-results").style.display = "none";
      document.getElementById("zodiac-results").style.display = "block";
      document.getElementById("numerology-results").style.display = "none";
      fetch(
        "/zodiac?dob=" +
          dob +
          "&firstName=" +
          firstName +
          "&lastName=" +
          lastName
      )
        .then((response) => response.json())
        .then((data) => {
          var output;
          output = "<h3>Hey "+ firstName + ", you are a "+  data.zodiac_sign+".</h3> <i>Let’s look into your sign a little more.</i></br>" +  data.zodiac_sign_info +"<br/><br/>It seems that the sign you are the most compatible with is "+
          data.compatibility + 
          ". Very interesting… Your lucky number today is "+ data.lucky_number +
          " and your lucky time is "+ data.lucky_time + ", how lucky you are. It also says here that your favourite color is "+data.color +", great choice!" +
          "</br> </br>The zodiac is a diagram used by astrologers to represent the positions of the planets and stars. It is divided into twelve sections, each of which has its own name and symbol. The zodiac is used to try to calculate the influence of the planets, especially on someone's life. (Source: https://www.collinsdictionary.com/)";
          document.getElementById("zodiac-output").innerHTML = output;
          location.hash = "#zodiac-results";

            //function that changes the hash
            function setHash(newHash) {
                location.hash = 'someHashThatDoesntExist';
                location.hash = newHash;
            }

            setHash('#zodiac-results');
        });
    }
  }

  //Numerology
  function getNumerologySection() {
    var dob = document.getElementById("dob").value;
    var firstName = document.getElementById("firstName").value;
    var lastName = document.getElementById("lastName").value;

    var today = new Date();
    var dd = String(today.getDate()).padStart(2, "0");
    var mm = String(today.getMonth() + 1).padStart(2, "0");
    var yyyy = today.getFullYear();
    today = yyyy + "-" + mm + "-" + dd;

    if (dob >= today) {
      alertDob.style.display = "block";
    } else if (firstName.trim() === "" || firstName.match(/\d+/g)!= null) {
      alertFName.style.display = "block";
    } else if (lastName.trim() === "" || lastName.match(/\d+/g)!= null) {
      alertLName.style.display = "block";
    } else if (form.checkValidity() == false) {
      var list = form.querySelectorAll(":invalid");
      for (var item of list) {
        item.focus();
        alertDob.style.display = "block";
      }
    } else {
      alertFName.style.display = "none";
      alertLName.style.display = "none";
      alertDob.style.display = "none";
      document.getElementById("funfact-results").style.display = "none";
      document.getElementById("zodiac-results").style.display = "none";
      document.getElementById("numerology-results").style.display = "block";
      fetch(
        "/numerology?dob=" +
          dob +
          "&firstName=" +
          firstName +
          "&lastName=" +
          lastName
      )
        .then((response) => response.json())
        .then((data) => {
          var output;
          output ="<h3>Hey "+ firstName +", Your life path number is "+  data.life_path_number+".</h3>Let’s find out a little about what that means for you.</br>"+
          data.life_path_meaning + ".</br><br/><h3> Your destiny number is "+ data.destiny_number +
          ".</h3>This number means that "+data.destiny_meaning;
          document.getElementById("numerology-output").innerHTML = output;
          location.hash = "#numerology-results";
            //function that changes the hash
            function setHash(newHash) {
                location.hash = 'someHashThatDoesntExist';
                location.hash = newHash;
            }
            setHash('#numerology-results');
        });
    }
  }
