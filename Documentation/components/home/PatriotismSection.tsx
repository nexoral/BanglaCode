"use client";

import { Globe, Heart, Zap, Users } from "lucide-react";

const initiatives = [
  {
    icon: Heart,
    title: "Cultural Pride",
    description: "Celebrate Bengali heritage through technology. Programming in your native language is not just functional—it's a statement of cultural pride.",
    color: "text-red-500",
    bgColor: "bg-red-500/10",
  },
  {
    icon: Globe,
    title: "Inclusive Technology",
    description: "Tech shouldn't be a Western monopoly. BanglaCode brings 300 million Bengali speakers into the digital revolution on their own terms.",
    color: "text-blue-500",
    bgColor: "bg-blue-500/10",
  },
  {
    icon: Zap,
    title: "Empowerment",
    description: "Remove language barriers. Unleash the potential of brilliant minds who think in Bengali but are forced to code in English.",
    color: "text-yellow-500",
    bgColor: "bg-yellow-500/10",
  },
  {
    icon: Users,
    title: "Community Building",
    description: "Create a community of Bengali developers, entrepreneurs, and educators who contribute to India's tech ecosystem.",
    color: "text-green-500",
    bgColor: "bg-green-500/10",
  },
];

export default function PatriotismSection() {
  return (
    <section className="py-32 bg-gradient-to-b from-background via-accent/5 to-background relative overflow-hidden">
      {/* Background Elements */}
      <div className="absolute inset-0 opacity-20">
        <div className="absolute top-1/3 left-0 w-96 h-96 bg-orange-500/10 rounded-full blur-3xl" />
        <div className="absolute bottom-1/3 right-0 w-96 h-96 bg-green-500/10 rounded-full blur-3xl" />
      </div>

      <div className="container mx-auto px-4 relative z-10">
        {/* Header */}
        <div className="max-w-4xl mx-auto text-center mb-20 animate-fade-in-up">
          <div className="inline-flex items-center gap-2 px-4 py-2 text-sm font-semibold text-orange-500 bg-orange-500/10 rounded-full border border-orange-500/20 mb-6">
            <span className="text-2xl">🇮🇳</span>
            Made in India, For the World
          </div>

          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            More Than Code
            <br />
            <span className="bg-clip-text text-transparent bg-gradient-to-r from-orange-500 via-red-500 to-green-500">
              A Movement for Indian Tech
            </span>
          </h2>

          <p className="text-xl text-muted-foreground max-w-3xl mx-auto leading-relaxed">
            BanglaCode is not just another programming language. It's a statement that technology belongs to everyone, regardless of their mother tongue. It's a celebration of Indian culture in the digital age, and a bridge that connects centuries of linguistic tradition with cutting-edge technology.
          </p>
        </div>

        {/* Grid of Initiatives */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mb-20">
          {initiatives.map((item, index) => (
            <div
              key={index}
              className="group relative animate-fade-in-up"
              style={{ animationDelay: `${index * 0.1}s` }}
            >
              {/* Card */}
              <div className="relative bg-card border border-border rounded-2xl p-8 hover:border-primary/50 hover:shadow-xl hover:-translate-y-1 transition-all duration-200">
                {/* Icon */}
                <div
                  className={`w-14 h-14 rounded-xl ${item.bgColor} flex items-center justify-center mb-6 group-hover:scale-110 transition-transform duration-200`}
                >
                  <item.icon className={`w-7 h-7 ${item.color}`} />
                </div>

                {/* Content */}
                <h3 className="text-2xl font-bold mb-4">{item.title}</h3>
                <p className="text-muted-foreground leading-relaxed">{item.description}</p>

                {/* Accent Line */}
                <div className={`absolute inset-x-0 bottom-0 h-1 ${item.bgColor} rounded-full scale-x-0 group-hover:scale-x-100 transition-transform duration-300 origin-left`} />
              </div>
            </div>
          ))}
        </div>

        {/* Bengali Philosophy Section */}
        <div className="max-w-4xl mx-auto mb-20 animate-fade-in-up" style={{ animationDelay: "0.4s" }}>
          <div className="relative">
            <div className="absolute inset-0 bg-gradient-to-r from-orange-500/20 via-pink-500/20 to-green-500/20 rounded-3xl blur-2xl" />
            <div className="relative bg-card border border-border/50 rounded-3xl p-12 backdrop-blur-sm">
              <div className="text-center space-y-6">
                <h3 className="text-3xl font-bold">
                  আমাদের দর্শন
                  <br />
                  <span className="text-muted-foreground text-2xl mt-4 block">Our Philosophy</span>
                </h3>

                <blockquote className="text-xl md:text-2xl font-light italic text-muted-foreground">
                  "আমরা বিশ্বাস করি যে প্রযুক্তি আমাদের ভাষায় কথা বলতে পারে। বাংলায় কোড লেখা শুধু একটি ভাষা নয়, এটি একটি আন্দোলন।"
                </blockquote>

                <p className="text-lg text-muted-foreground">
                  "We believe technology can speak our language. Coding in Bengali is not just about language—it's a movement for cultural representation in the digital world."
                </p>

                <p className="text-primary font-semibold text-lg">
                  — BanglaCode Creators, West Bengal, India 🇮🇳
                </p>
              </div>
            </div>
          </div>
        </div>

        {/* Impact Metrics */}
        <div className="max-w-5xl mx-auto animate-fade-in-up" style={{ animationDelay: "0.5s" }}>
          <h3 className="text-3xl font-bold text-center mb-12">Our Vision</h3>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {/* Metric 1 */}
            <div className="text-center group hover:bg-accent/50 p-8 rounded-2xl transition-all duration-200">
              <div className="text-5xl font-black bg-clip-text text-transparent bg-gradient-to-r from-orange-500 to-red-500 mb-4">
                300M+
              </div>
              <p className="text-lg font-semibold mb-2">Bengali Speakers</p>
              <p className="text-muted-foreground">Potential developers with access to native-language programming</p>
            </div>

            {/* Metric 2 */}
            <div className="text-center group hover:bg-accent/50 p-8 rounded-2xl transition-all duration-200">
              <div className="text-5xl font-black bg-clip-text text-transparent bg-gradient-to-r from-green-500 to-emerald-500 mb-4">
                Zero
              </div>
              <p className="text-lg font-semibold mb-2">Language Barriers</p>
              <p className="text-muted-foreground">Think in Bengali. Code in Bengali. No translation needed.</p>
            </div>

            {/* Metric 3 */}
            <div className="text-center group hover:bg-accent/50 p-8 rounded-2xl transition-all duration-200">
              <div className="text-5xl font-black bg-clip-text text-transparent bg-gradient-to-r from-blue-500 to-cyan-500 mb-4">
                100%
              </div>
              <p className="text-lg font-semibold mb-2">Open Source</p>
              <p className="text-muted-foreground">Free, transparent, community-driven development for everyone</p>
            </div>
          </div>
        </div>

        {/* Call to Action */}
        <div className="max-w-3xl mx-auto mt-20 text-center animate-fade-in-up" style={{ animationDelay: "0.6s" }}>
          <div className="bg-gradient-to-r from-orange-500/10 via-red-500/10 to-green-500/10 border border-red-500/20 rounded-2xl p-10">
            <h3 className="text-2xl md:text-3xl font-bold mb-6">
              Join the Bengali Tech Revolution
            </h3>
            <p className="text-lg text-muted-foreground mb-8">
              BanglaCode is more than a language—it's a symbol of Indian excellence in technology. Whether you're a student, educator, or developer, you're part of something bigger: proving that Indians can lead global tech innovation on their own terms.
            </p>
            <div className="flex flex-wrap gap-4 justify-center text-sm font-semibold">
              <span className="px-4 py-2 bg-orange-500/20 rounded-full text-orange-400 border border-orange-500/30">
                🇮🇳 Indian-Made
              </span>
              <span className="px-4 py-2 bg-green-500/20 rounded-full text-green-400 border border-green-500/30">
                🗣️ Bengali-First
              </span>
              <span className="px-4 py-2 bg-blue-500/20 rounded-full text-blue-400 border border-blue-500/30">
                🌍 Globally Accessible
              </span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
