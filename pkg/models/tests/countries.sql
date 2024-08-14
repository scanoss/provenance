DROP TABLE if exists countries;

CREATE TABLE countries (
	country_name text NOT NULL,
	id int4,
	CONSTRAINT countries_pkey PRIMARY KEY (id),
	CONSTRAINT countries_unique UNIQUE (country_name)
);

INSERT INTO countries
(country_name, id)
VALUES('Afghanistan', 1);
INSERT INTO countries
(country_name, id)
VALUES('Albania', 2);
INSERT INTO countries
(country_name, id)
VALUES('Algeria', 3);
INSERT INTO countries
(country_name, id)
VALUES('Andorra', 4);
INSERT INTO countries
(country_name, id)
VALUES('Angola', 5);
INSERT INTO countries
(country_name, id)
VALUES('Antigua and Barbuda', 6);
INSERT INTO countries
(country_name, id)
VALUES('Argentina', 7);
INSERT INTO countries
(country_name, id)
VALUES('Armenia', 8);
INSERT INTO countries
(country_name, id)
VALUES('Australia', 9);
INSERT INTO countries
(country_name, id)
VALUES('Austria', 10);
INSERT INTO countries
(country_name, id)
VALUES('Azerbaijan', 11);
INSERT INTO countries
(country_name, id)
VALUES('Bahamas', 12);
INSERT INTO countries
(country_name, id)
VALUES('Bahrain', 13);
INSERT INTO countries
(country_name, id)
VALUES('Bangladesh', 14);
INSERT INTO countries
(country_name, id)
VALUES('Barbados', 15);
INSERT INTO countries
(country_name, id)
VALUES('Belarus', 16);
INSERT INTO countries
(country_name, id)
VALUES('Belgium', 17);
INSERT INTO countries
(country_name, id)
VALUES('Belize', 18);
INSERT INTO countries
(country_name, id)
VALUES('Benin', 19);
INSERT INTO countries
(country_name, id)
VALUES('Bhutan', 20);
INSERT INTO countries
(country_name, id)
VALUES('Bolivia', 21);
INSERT INTO countries
(country_name, id)
VALUES('Bosnia and Herzegovina', 22);
INSERT INTO countries
(country_name, id)
VALUES('Botswana', 23);
INSERT INTO countries
(country_name, id)
VALUES('Brazil', 24);
INSERT INTO countries
(country_name, id)
VALUES('Brunei', 25);
INSERT INTO countries
(country_name, id)
VALUES('Bulgaria', 26);
INSERT INTO countries
(country_name, id)
VALUES('Burkina Faso', 27);
INSERT INTO countries
(country_name, id)
VALUES('Burundi', 28);
INSERT INTO countries
(country_name, id)
VALUES('Cabo Verde', 29);
INSERT INTO countries
(country_name, id)
VALUES('Cambodia', 30);
INSERT INTO countries
(country_name, id)
VALUES('Cameroon', 31);
INSERT INTO countries
(country_name, id)
VALUES('Canada', 32);
INSERT INTO countries
(country_name, id)
VALUES('Central African Republic', 33);
INSERT INTO countries
(country_name, id)
VALUES('Chad', 34);
INSERT INTO countries
(country_name, id)
VALUES('Chile', 35);
INSERT INTO countries
(country_name, id)
VALUES('China', 36);
INSERT INTO countries
(country_name, id)
VALUES('Colombia', 37);
INSERT INTO countries
(country_name, id)
VALUES('Comoros', 38);
INSERT INTO countries
(country_name, id)
VALUES('Congo, Democratic Republic of the', 39);
INSERT INTO countries
(country_name, id)
VALUES('Congo, Republic of the', 40);
INSERT INTO countries
(country_name, id)
VALUES('Costa Rica', 41);
INSERT INTO countries
(country_name, id)
VALUES('Croatia', 42);
INSERT INTO countries
(country_name, id)
VALUES('Cuba', 43);
INSERT INTO countries
(country_name, id)
VALUES('Cyprus', 44);
INSERT INTO countries
(country_name, id)
VALUES('Czech Republic', 45);
INSERT INTO countries
(country_name, id)
VALUES('Denmark', 46);
INSERT INTO countries
(country_name, id)
VALUES('Djibouti', 47);
INSERT INTO countries
(country_name, id)
VALUES('Dominica', 48);
INSERT INTO countries
(country_name, id)
VALUES('Dominican Republic', 49);
INSERT INTO countries
(country_name, id)
VALUES('East Timor (Timor-Leste)', 50);
INSERT INTO countries
(country_name, id)
VALUES('Ecuador', 51);
INSERT INTO countries
(country_name, id)
VALUES('Egypt', 52);
INSERT INTO countries
(country_name, id)
VALUES('El Salvador', 53);
INSERT INTO countries
(country_name, id)
VALUES('Equatorial Guinea', 54);
INSERT INTO countries
(country_name, id)
VALUES('Eritrea', 55);
INSERT INTO countries
(country_name, id)
VALUES('Estonia', 56);
INSERT INTO countries
(country_name, id)
VALUES('Ethiopia', 57);
INSERT INTO countries
(country_name, id)
VALUES('Fiji', 58);
INSERT INTO countries
(country_name, id)
VALUES('Finland', 59);
INSERT INTO countries
(country_name, id)
VALUES('France', 60);
INSERT INTO countries
(country_name, id)
VALUES('Gabon', 61);
INSERT INTO countries
(country_name, id)
VALUES('Gambia', 62);
INSERT INTO countries
(country_name, id)
VALUES('Georgia', 63);
INSERT INTO countries
(country_name, id)
VALUES('Germany', 64);
INSERT INTO countries
(country_name, id)
VALUES('Ghana', 65);
INSERT INTO countries
(country_name, id)
VALUES('Greece', 66);
INSERT INTO countries
(country_name, id)
VALUES('Grenada', 67);
INSERT INTO countries
(country_name, id)
VALUES('Guatemala', 68);
INSERT INTO countries
(country_name, id)
VALUES('Guinea', 69);
INSERT INTO countries
(country_name, id)
VALUES('Guinea-Bissau', 70);
INSERT INTO countries
(country_name, id)
VALUES('Guyana', 71);
INSERT INTO countries
(country_name, id)
VALUES('Haiti', 72);
INSERT INTO countries
(country_name, id)
VALUES('Honduras', 73);
INSERT INTO countries
(country_name, id)
VALUES('Hungary', 74);
INSERT INTO countries
(country_name, id)
VALUES('Iceland', 75);
INSERT INTO countries
(country_name, id)
VALUES('India', 76);
INSERT INTO countries
(country_name, id)
VALUES('Indonesia', 77);
INSERT INTO countries
(country_name, id)
VALUES('Iran', 78);
INSERT INTO countries
(country_name, id)
VALUES('Iraq', 79);
INSERT INTO countries
(country_name, id)
VALUES('Ireland', 80);
INSERT INTO countries
(country_name, id)
VALUES('Israel', 81);
INSERT INTO countries
(country_name, id)
VALUES('Italy', 82);
INSERT INTO countries
(country_name, id)
VALUES('Jamaica', 83);
INSERT INTO countries
(country_name, id)
VALUES('Japan', 84);
INSERT INTO countries
(country_name, id)
VALUES('Jordan', 85);
INSERT INTO countries
(country_name, id)
VALUES('Kazakhstan', 86);
INSERT INTO countries
(country_name, id)
VALUES('Kenya', 87);
INSERT INTO countries
(country_name, id)
VALUES('Kiribati', 88);
INSERT INTO countries
(country_name, id)
VALUES('Korea, North', 89);
INSERT INTO countries
(country_name, id)
VALUES('Korea, South', 90);
INSERT INTO countries
(country_name, id)
VALUES('Kosovo', 91);
INSERT INTO countries
(country_name, id)
VALUES('Kuwait', 92);
INSERT INTO countries
(country_name, id)
VALUES('Kyrgyzstan', 93);
INSERT INTO countries
(country_name, id)
VALUES('Laos', 94);
INSERT INTO countries
(country_name, id)
VALUES('Latvia', 95);
INSERT INTO countries
(country_name, id)
VALUES('Lebanon', 96);
INSERT INTO countries
(country_name, id)
VALUES('Lesotho', 97);
INSERT INTO countries
(country_name, id)
VALUES('Liberia', 98);
INSERT INTO countries
(country_name, id)
VALUES('Libya', 99);
INSERT INTO countries
(country_name, id)
VALUES('Liechtenstein', 100);
INSERT INTO countries
(country_name, id)
VALUES('Lithuania', 101);
INSERT INTO countries
(country_name, id)
VALUES('Luxembourg', 102);
INSERT INTO countries
(country_name, id)
VALUES('Macedonia', 103);
INSERT INTO countries
(country_name, id)
VALUES('Madagascar', 104);
INSERT INTO countries
(country_name, id)
VALUES('Malawi', 105);
INSERT INTO countries
(country_name, id)
VALUES('Malaysia', 106);
INSERT INTO countries
(country_name, id)
VALUES('Maldives', 107);
INSERT INTO countries
(country_name, id)
VALUES('Mali', 108);
INSERT INTO countries
(country_name, id)
VALUES('Malta', 109);
INSERT INTO countries
(country_name, id)
VALUES('Marshall Islands', 110);
INSERT INTO countries
(country_name, id)
VALUES('Mauritania', 111);
INSERT INTO countries
(country_name, id)
VALUES('Mauritius', 112);
INSERT INTO countries
(country_name, id)
VALUES('Mexico', 113);
INSERT INTO countries
(country_name, id)
VALUES('Micronesia', 114);
INSERT INTO countries
(country_name, id)
VALUES('Moldova', 115);
INSERT INTO countries
(country_name, id)
VALUES('Monaco', 116);
INSERT INTO countries
(country_name, id)
VALUES('Mongolia', 117);
INSERT INTO countries
(country_name, id)
VALUES('Montenegro', 118);
INSERT INTO countries
(country_name, id)
VALUES('Morocco', 119);
INSERT INTO countries
(country_name, id)
VALUES('Mozambique', 120);
INSERT INTO countries
(country_name, id)
VALUES('Myanmar', 121);
INSERT INTO countries
(country_name, id)
VALUES('Namibia', 122);
INSERT INTO countries
(country_name, id)
VALUES('Nauru', 123);
INSERT INTO countries
(country_name, id)
VALUES('Nepal', 124);
INSERT INTO countries
(country_name, id)
VALUES('Netherlands', 125);
INSERT INTO countries
(country_name, id)
VALUES('New Zealand', 126);
INSERT INTO countries
(country_name, id)
VALUES('Nicaragua', 127);
INSERT INTO countries
(country_name, id)
VALUES('Niger', 128);
INSERT INTO countries
(country_name, id)
VALUES('Nigeria', 129);
INSERT INTO countries
(country_name, id)
VALUES('Norway', 130);
INSERT INTO countries
(country_name, id)
VALUES('Oman', 131);
INSERT INTO countries
(country_name, id)
VALUES('Pakistan', 132);
INSERT INTO countries
(country_name, id)
VALUES('Palau', 133);
INSERT INTO countries
(country_name, id)
VALUES('Panama', 134);
INSERT INTO countries
(country_name, id)
VALUES('Papua New Guinea', 135);
INSERT INTO countries
(country_name, id)
VALUES('Paraguay', 136);
INSERT INTO countries
(country_name, id)
VALUES('Peru', 137);
INSERT INTO countries
(country_name, id)
VALUES('Philippines', 138);
INSERT INTO countries
(country_name, id)
VALUES('Poland', 139);
INSERT INTO countries
(country_name, id)
VALUES('Portugal', 140);
INSERT INTO countries
(country_name, id)
VALUES('Qatar', 141);
INSERT INTO countries
(country_name, id)
VALUES('Romania', 142);
INSERT INTO countries
(country_name, id)
VALUES('Russia', 143);
INSERT INTO countries
(country_name, id)
VALUES('Rwanda', 144);
INSERT INTO countries
(country_name, id)
VALUES('Saint Kitts and Nevis', 145);
INSERT INTO countries
(country_name, id)
VALUES('Saint Lucia', 146);
INSERT INTO countries
(country_name, id)
VALUES('Saint Vincent and the Grenadines', 147);
INSERT INTO countries
(country_name, id)
VALUES('Samoa', 148);
INSERT INTO countries
(country_name, id)
VALUES('San Marino', 149);
INSERT INTO countries
(country_name, id)
VALUES('Sao Tome and Principe', 150);
INSERT INTO countries
(country_name, id)
VALUES('Saudi Arabia', 151);
INSERT INTO countries
(country_name, id)
VALUES('Senegal', 152);
INSERT INTO countries
(country_name, id)
VALUES('Serbia', 153);
INSERT INTO countries
(country_name, id)
VALUES('Seychelles', 154);
INSERT INTO countries
(country_name, id)
VALUES('Sierra Leone', 155);
INSERT INTO countries
(country_name, id)
VALUES('Singapore', 156);
INSERT INTO countries
(country_name, id)
VALUES('Slovakia', 157);
INSERT INTO countries
(country_name, id)
VALUES('Slovenia', 158);
INSERT INTO countries
(country_name, id)
VALUES('Solomon Islands', 159);
INSERT INTO countries
(country_name, id)
VALUES('Somalia', 160);
INSERT INTO countries
(country_name, id)
VALUES('South Africa', 161);
INSERT INTO countries
(country_name, id)
VALUES('South Sudan', 162);
INSERT INTO countries
(country_name, id)
VALUES('Spain', 163);
INSERT INTO countries
(country_name, id)
VALUES('Sri Lanka', 164);
INSERT INTO countries
(country_name, id)
VALUES('Sudan', 165);
INSERT INTO countries
(country_name, id)
VALUES('Suriname', 166);
INSERT INTO countries
(country_name, id)
VALUES('Swaziland', 167);
INSERT INTO countries
(country_name, id)
VALUES('Sweden', 168);
INSERT INTO countries
(country_name, id)
VALUES('Switzerland', 169);
INSERT INTO countries
(country_name, id)
VALUES('Syria', 170);
INSERT INTO countries
(country_name, id)
VALUES('Taiwan', 171);
INSERT INTO countries
(country_name, id)
VALUES('Tajikistan', 172);
INSERT INTO countries
(country_name, id)
VALUES('Tanzania', 173);
INSERT INTO countries
(country_name, id)
VALUES('Thailand', 174);
INSERT INTO countries
(country_name, id)
VALUES('Togo', 175);
INSERT INTO countries
(country_name, id)
VALUES('Tonga', 176);
INSERT INTO countries
(country_name, id)
VALUES('Trinidad and Tobago', 177);
INSERT INTO countries
(country_name, id)
VALUES('Tunisia', 178);
INSERT INTO countries
(country_name, id)
VALUES('Turkey', 179);
INSERT INTO countries
(country_name, id)
VALUES('Turkmenistan', 180);
INSERT INTO countries
(country_name, id)
VALUES('Tuvalu', 181);
INSERT INTO countries
(country_name, id)
VALUES('Uganda', 182);
INSERT INTO countries
(country_name, id)
VALUES('Ukraine', 183);
INSERT INTO countries
(country_name, id)
VALUES('United Arab Emirates', 184);
INSERT INTO countries
(country_name, id)
VALUES('United Kingdom', 185);
INSERT INTO countries
(country_name, id)
VALUES('United States', 186);
INSERT INTO countries
(country_name, id)
VALUES('Uruguay', 187);
INSERT INTO countries
(country_name, id)
VALUES('Uzbekistan', 188);
INSERT INTO countries
(country_name, id)
VALUES('Vanuatu', 189);
INSERT INTO countries
(country_name, id)
VALUES('Vatican City', 190);
INSERT INTO countries
(country_name, id)
VALUES('Venezuela', 191);
INSERT INTO countries
(country_name, id)
VALUES('Vietnam', 192);
INSERT INTO countries
(country_name, id)
VALUES('Yemen', 193);
INSERT INTO countries
(country_name, id)
VALUES('Zambia', 194);
INSERT INTO countries
(country_name, id)
VALUES('Zimbabwe', 195);
INSERT INTO countries
(country_name, id)
VALUES('State of Palestine', 196);
INSERT INTO countries
(country_name, id)
VALUES('new_test1', 197);
INSERT INTO countries
(country_name, id)
VALUES('new_test2', 198);
INSERT INTO countries
(country_name, id)
VALUES('new_test3', 199);
INSERT INTO countries
(country_name, id)
VALUES('new_test4', 201);
