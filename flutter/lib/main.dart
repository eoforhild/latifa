import 'dart:convert';
import 'package:file_picker/file_picker.dart';

import 'package:flutter/material.dart';
import 'ffi.dart' if (dart.library.html) 'ffi_web.dart';

void main() {
  api.generateKeysAndDump();
  runApp(MyApp());
}

class RegisterScreen extends StatelessWidget {
  final passwordController = TextEditingController();
  final usernameController = TextEditingController();
  final emailController = TextEditingController();

  RegisterScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Register Screen'),
      ),
      body: Center(
          child: Form(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
              child: TextFormField(
                controller: usernameController,
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Enter your username',
                ),
              ),
            ),
            const SizedBox(height: 10),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
              child: TextFormField(
                controller: passwordController,
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Enter your Password',
                ),
              ),
            ),
            const SizedBox(height: 10),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
              child: TextFormField(
                controller: emailController,
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Enter your Email',
                ),
              ),
            ),
            const SizedBox(height: 10),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
              child: ElevatedButton(
                // Within the `FirstScreen` widget
                onPressed: () {
                  Map preJson = {
                    "username": usernameController.text,
                    "email": emailController.text,
                    "password": passwordController.text,
                  };
                  String stringJson = json.encode(preJson);
                  api.registerAndPublish(regForm: stringJson);

                  Navigator.pushNamed(context, '/login');
                },
                child: const Text('Submit'),
              ),
            ),
          ],
        ),
      )),
    );
  }
}

class LoginScreen extends StatelessWidget {
  final passwordController = TextEditingController();
  final usernameController = TextEditingController();

  LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text('Login Screen'),
        ),
        body: Center(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
                child: TextFormField(
                  controller: usernameController,
                  decoration: const InputDecoration(
                    border: UnderlineInputBorder(),
                    labelText: 'Enter your Email',
                  ),
                ),
              ),
              const SizedBox(height: 10),
              Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
                child: TextFormField(
                  controller: passwordController,
                  decoration: const InputDecoration(
                    border: UnderlineInputBorder(),
                    labelText: 'Enter your Password',
                  ),
                ),
              ),
              const SizedBox(height: 10),
              Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 40, vertical: 16),
                child: ElevatedButton(
                  // Within the `FirstScreen` widget
                  onPressed: () async {
                    Map preJson = {
                      "login": usernameController.text,
                      "password": passwordController.text
                    };
                    String stringJson = json.encode(preJson);
                    if (await api.login(logForm: stringJson)) {
                      Navigator.pushNamed(context, '/main');
                    }
                  },
                  child: const Text('Submit'),
                ),
              ),
            ],
          ),
        ));
  }
}

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Home Screen'),
      ),
      body: Center(
          child: Wrap(
        crossAxisAlignment: WrapCrossAlignment.center,
        direction: Axis.vertical,
        spacing: 20,
        children: <Widget>[
          ElevatedButton(
            // Within the `FirstScreen` widget
            onPressed: () {
              // Navigate to the second screen using a named route.
              Navigator.pushNamed(context, '/login');
            },
            child: const Text('Login'),
          ),
          ElevatedButton(
            // Within the `FirstScreen` widget
            onPressed: () {
              // Navigate to the second screen using a named route.
              Navigator.pushNamed(context, '/register');
            },
            child: const Text('Register'),
          ),
        ],
      )),
    );
  }
}

class MainScreen extends StatefulWidget {
  const MainScreen({super.key});

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  int currentPageIndex = 0;
  final emailController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: NavigationBar(
        onDestinationSelected: (int index) {
          setState(() {
            currentPageIndex = index;
          });
        },
        indicatorColor: Colors.amber[800],
        selectedIndex: currentPageIndex,
        destinations: const <Widget>[
          NavigationDestination(
            selectedIcon: Icon(Icons.home),
            icon: Icon(Icons.home_outlined),
            label: 'Home',
          ),
          NavigationDestination(
            icon: Icon(Icons.business),
            label: 'Business',
          ),
          NavigationDestination(
            selectedIcon: Icon(Icons.school),
            icon: Icon(Icons.school_outlined),
            label: 'School',
          ),
        ],
      ),
      body: <Widget>[
        Scaffold(
            appBar: AppBar(
              title: Text('LATIFA'),
              automaticallyImplyLeading: false,
              actions: [
                IconButton(
                  icon: const Icon(Icons.logout),
                  tooltip: 'Logout',
                  onPressed: () {
                    api.logout();
                    Navigator.pushNamed(context, '/');
                  },
                ),
              ],
            ),
            body: Center(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Padding(
                    padding: const EdgeInsets.symmetric(
                        horizontal: 40, vertical: 16),
                    child: TextFormField(
                      controller: emailController,
                      decoration: const InputDecoration(
                        border: UnderlineInputBorder(),
                        labelText: 'Enter the email of your intended recipient',
                      ),
                    ),
                  ),
                  const SizedBox(height: 10),
                  ElevatedButton(
                      child: Text('Select file and send!'),
                      onPressed: () async {
                        var picked = await FilePicker.platform.pickFiles();
                        var email = emailController.text;

                        Navigator.pushNamed(context, '/main');
                      }),
                ],
              ),
            )),
        Scaffold(
          appBar: AppBar(
            title: Text('LATIFA'),
            automaticallyImplyLeading: false,
            actions: [
              IconButton(
                icon: const Icon(Icons.logout),
                tooltip: 'Logout',
                onPressed: () {
                  api.logout();
                  Navigator.pushNamed(context, '/');
                },
              ),
            ],
          ),
          body: Container(
            color: Colors.green,
            alignment: Alignment.center,
            child: const Text('Page 2'),
          ),
        ),
        Scaffold(
          appBar: AppBar(
            title: Text('LATIFA'),
            automaticallyImplyLeading: false,
            actions: [
              IconButton(
                icon: const Icon(Icons.logout),
                tooltip: 'Logout',
                onPressed: () {
                  api.logout();
                  Navigator.pushNamed(context, '/');
                },
              ),
            ],
          ),
          body: Container(
            color: Colors.blue,
            alignment: Alignment.center,
            child: const Text('Page 3'),
          ),
        ),
      ][currentPageIndex],
    );
  }
}

class FileUploadButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      child: Text('UPLOAD FILE'),
      onPressed: () async {
        var picked = await FilePicker.platform.pickFiles();
      },
    );
  }
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.green,
        useMaterial3: true,
      ),
      //home: const MyHomePage(title: 'Flutter Demo Home Page'),
      initialRoute: '/',
      routes: {
        // When navigating to the "/" route, build the FirstScreen widget.
        '/': (context) => HomeScreen(),
        // When navigating to the "/second" route, build the SecondScreen widget.
        '/login': (context) => LoginScreen(),
        '/register': (context) => RegisterScreen(),
        '/main': (context) => MainScreen(),
      },
    );
  }
}
